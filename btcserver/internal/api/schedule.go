package api

import (
	"strconv"
	"sync"
	"time"

	"github.com/generativelabs/btcserver/internal/chakra"
	"github.com/generativelabs/btcserver/internal/types"
	"github.com/generativelabs/btcserver/internal/utils"
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

		v, _ := strconv.Atoi(tw.Value)

		seconds := v / 1000
		nanos := (v % 1000) * 1000000
		t := time.Unix(int64(seconds), int64(nanos)).UTC()
		s.ScheduleTimeWheel = t
	} else {
		s.ScheduleTimeWheel = time.Now().UTC()
		err = s.backend.CreateTimeWheel(s.ScheduleTimeWheel.UnixMilli())
		if err != nil {
			log.Error().Msgf("❌ error create time wheel exist: %s ", err)
			return
		}
	}

	for {
		// Start from the global timestamp, query every 5 minutes, and query all users to be released within 5 minutes
		txs, err := s.backend.QueryAllNotYetLockedUpTxNextPeriod(s.ScheduleTimeWheel.UnixMilli())
		if err != nil {
			log.Error().Msgf("❌ error query all not release tx: %s ", err)
			time.Sleep(time.Second)
			continue
		}

		if len(txs) == 0 {
			log.Info().Msgf("No tx to be released, sleep %v minutes", s.ScheduleTimeWheel.Add(5*time.Minute).Sub(time.Now()).Minutes())
			err = s.UpdateTimeWheel()
			if err != nil {
				log.Error().Msgf("❌ error update time wheel: %s ", err)
			}
			time.Sleep(time.Until(s.ScheduleTimeWheel.Add(5 * time.Minute)))
			continue
		}

		go s.RewardTasksSchedule(txs)
	}
}

func (s *Server) RewardTasksSchedule(txs []*types.ReleaseTxsInfo) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(tx []*types.ReleaseTxsInfo) {
		defer wg.Done()
		err := s.RewardTasks(tx)
		if err != nil {
			log.Error().Msgf("❌ error reward tasks: %s ", err)
			return
		}
	}(txs)

	wg.Wait()

	_ = s.UpdateTimeWheel()
}

func (s *Server) RewardTasks(txs []*types.ReleaseTxsInfo) error {
	seconds := txs[0].ReleasingTime / 1000
	nanos := (txs[0].ReleasingTime % 1000) * 1000000
	t := time.Unix(seconds, nanos).UTC()

	timer := time.NewTimer(time.Until(t))

	for { //nolint
		select {
		case <-timer.C:
			txIDs := make([]string, 0)
			for _, tx := range txs {
				txIDs = append(txIDs, tx.Tx)
			}

			res, err := chakra.RewardTo(s.Ctx, s.ChakraAccount, s.ContractAddress, txIDs)
			if err != nil {
				log.Error().Msgf("❌ error reward to txIDs %v: %s ", txIDs, err)
				// todo deal err task
				return err
			}

			log.Info().Msgf("Chakra reward to success, txs hash: %s ", res.TransactionHash)
			for _, tx := range txs {
				err = s.backend.UpdateStakeReleasingTime(tx.Staker, tx.Tx)
				if err != nil {
					log.Error().Msgf("❌ error %s update stake releasing time to txID %s: %s ", tx.Staker, tx.Tx, err)
					return err
				}
			}

			return nil
		}
	}
}

func (s *Server) UpdateTimeWheel() error {
	newScheduleTimeWheel := utils.TimeTOTimestamp(s.ScheduleTimeWheel) + 5*time.Minute.Milliseconds()
	log.Info().Msgf("Update time wheel to %s ", utils.TimestampToTime(newScheduleTimeWheel))
	err := s.backend.UpdateTimeWheel(newScheduleTimeWheel)
	if err != nil {
		log.Error().Msgf("❌ error update time wheel for db: %s ", err)
		return err
	}

	s.ScheduleTimeWheel = s.ScheduleTimeWheel.Add(5 * time.Minute)
	return nil
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
