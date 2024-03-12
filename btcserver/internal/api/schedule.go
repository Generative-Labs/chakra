package api

import (
	"sync"
	"time"

	"github.com/generativelabs/btcserver/internal"
	"github.com/generativelabs/btcserver/internal/chakra"
	"github.com/rs/zerolog/log"
)

func (s *Server) TimeWheelSchedule() {
	exist, err := s.backend.IsTimeWheelExist()
	if err != nil {
		log.Error().Msgf("❌ error query time wheel exist: %s ", err)
		return
	}

	if exist {
		tw, err := s.backend.GetTimeWheel()
		if err != nil {
			log.Error().Msgf("❌ error query time wheel: %s ", err)
			return
		}

		seconds := tw.Value / 1000
		nanos := (tw.Value % 1000) * 1000000
		t := time.Unix(int64(seconds), int64(nanos)).UTC()
		s.ScheduleTimeWheel = t
	}

	for {
		// Start from the global timestamp, query every 5 minutes, and query all users to be released within 5 minutes
		txs, err := s.backend.QueryAllNotYetLockedUpTxNextFourHours(uint64(s.ScheduleTimeWheel.UnixMilli()))
		if err != nil {
			log.Error().Msgf("❌ error query all not release tx: %s ", err)
			time.Sleep(time.Second)
			continue
		}

		go s.RewardTasksSchedule(txs)
	}
}

func (s *Server) RewardTasksSchedule(txs []*internal.ReleaseTxsInfo) {
	wg := sync.WaitGroup{}
	wg.Add(len(txs))
	for _, tx := range txs {
		go func(tx *internal.ReleaseTxsInfo) {
			defer wg.Done()
			err := s.RewardTasks(tx)
			if err != nil {
				log.Error().Msgf("❌ error reward tasks: %s ", err)
				return
			}
		}(tx)
	}

	wg.Wait()

	newScheduleTimeWheel := uint64(s.ScheduleTimeWheel.UnixMilli() + 5*time.Minute.Milliseconds())
	err := s.backend.UpdateTimeWheel(newScheduleTimeWheel)
	if err != nil {
		log.Error().Msgf("❌ error update time wheel for db: %s ", err)
		return
	}
	s.ScheduleTimeWheel = s.ScheduleTimeWheel.Add(5 * time.Minute)
}

func (s *Server) RewardTasks(tx *internal.ReleaseTxsInfo) error {
	seconds := tx.ReleasingTime / 1000
	nanos := (tx.ReleasingTime % 1000) * 1000000
	t := time.Unix(int64(seconds), int64(nanos)).UTC()

	timer := time.NewTimer(time.Until(t))

	for { //nolint
		select {
		case <-timer.C:
			res, err := chakra.RewardTo(s.Ctx, s.ChakraAccount, s.ContractAddress, tx.TxID)
			if err != nil {
				log.Error().Msgf("❌ error reward to txID %s: %s ", tx.TxID, err)
				// todo deal err task
				return err
			}
			log.Info().Msgf("Chakra reward to success, tx hash: %s ", res.TransactionHash)
			return nil
		}
	}
}

// UpdateStakeStatus defines the periodic process of reading stake records from the database,
// and updating the status of records in the database based on the status of transactions on the BTC chain.
func (s *Server) UpdateStakeStatus() {
	timer := time.NewTicker(5 * time.Minute)

	for range timer.C {
		stakeVerifyParams, err := s.backend.QueryNoFinalizedStakeTx()
		if err != nil {
			log.Error().Msgf("💥 error when query no finalized stake tx %s", err)
			continue
		}

		newStatuses, err := s.btcClient.CheckStakeRecords(stakeVerifyParams)
		if err != nil {
			log.Error().Msgf("💥 error when check state records %s", err)
			continue
		}

		for i, status := range newStatuses {
			if status == stakeVerifyParams[i].FinalizedStatus {
				continue
			}

			err := s.backend.UpdateStakeFinalizedStatus(stakeVerifyParams[i].StakerPubKey, stakeVerifyParams[i].TxID, int(status))
			if err != nil {
				log.Error().Msgf("💥 error when update state finalize status %s", err)
			}
		}
	}
}
