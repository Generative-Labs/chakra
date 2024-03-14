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

func InitTimeWheel() time.Time {
	currentTime := time.Now()
	minute := currentTime.Minute()
	second := currentTime.Second()

	if minute%5 != 0 {
		nearestFiveMultiple := minute - (minute % 5)
		minute = nearestFiveMultiple
		second = 0
	}

	nearestTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), minute, second, 0, currentTime.Location())

	if nearestTime.After(currentTime) {
		nearestTime = nearestTime.Add(-5 * time.Minute)
	}

	return nearestTime
}

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

		stw := utils.TimestampToTime(int64(v))
		s.ScheduleTimeWheel = stw
	} else {
		s.ScheduleTimeWheel = InitTimeWheel()
		err = s.backend.CreateTimeWheel(s.ScheduleTimeWheel.UnixNano())
		if err != nil {
			log.Error().Msgf("❌ error create time wheel exist: %s ", err)
			return
		}
	}

	log.Info().Msgf("🔵🔵🔵 Start time wheel schedule, time wheel: %s", s.ScheduleTimeWheel)

	for {
		// Start from the global timestamp, query every 5 minutes, and query all users to be released within 5 minutes
		txs, err := s.backend.QueryAllNotYetLockedUpTxNextPeriod(s.ScheduleTimeWheel.UnixNano(), types.TimeWheelSize)
		if err != nil {
			log.Error().Msgf("❌ error query all not release tx: %s ", err)
			time.Sleep(time.Second)
			continue
		}

		if len(txs) == 0 {
			if time.Now().After(s.ScheduleTimeWheel) {
				err = s.UpdateTimeWheel()
				if err != nil {
					log.Error().Msgf("❌ error update time wheel: %s ", err)
				}

				continue
			}

			log.Info().Msgf("🔵🔵🔵 No tx to be released, sleep %v minutes", s.ScheduleTimeWheel.Add(types.TimeWheelSize).Sub(time.Now()).Minutes()) //nolint
			err = s.UpdateTimeWheel()
			if err != nil {
				log.Error().Msgf("❌ error update time wheel: %s ", err)
			}
			time.Sleep(time.Until(s.ScheduleTimeWheel.Add(types.TimeWheelSize)))
			continue
		}

		log.Info().Msgf("🔵🔵🔵 In the next time period %s, find %d txids that are about to be released: %v", s.ScheduleTimeWheel.Add(types.TimeWheelSize), len(txs), txs)
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
	rt := utils.TimestampToTime(txs[0].ReleasingTime)
	timer := time.NewTimer(time.Until(rt))
	log.Info().Msgf("🔵🔵🔵 Start the timer and prepare to wait %v min", time.Until(rt).Minutes())

	for { //nolint
		select {
		case <-timer.C:
			txIDs := make([]string, 0)
			for _, tx := range txs {
				txIDs = append(txIDs, tx.Tx)
			}

			log.Info().Msgf("🔵🔵🔵 Start reward to %v ", txIDs)
			res, err := chakra.RewardTo(s.Ctx, s.ChakraAccount, s.ContractAddress, txIDs)
			if err != nil {
				log.Error().Msgf("❌ error reward to txIDs %v: %s ", txIDs, err)
				// todo deal err task
				return err
			}

			log.Info().Msgf("🔵🔵🔵 Chakra reward to success, txs hash: %s ", res.TransactionHash)
			for _, tx := range txs {
				log.Info().Msgf("🔵🔵🔵 Start update stake ReleasingTime %v ", tx)
				err := s.backend.UpdateStakeReleasingTime(tx.Staker, tx.Tx)
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
	newScheduleTimeWheel := utils.TimeTOTimestamp(s.ScheduleTimeWheel) + types.TimeWheelSize.Nanoseconds()
	log.Info().Msgf("Update time wheel to %s ", utils.TimestampToTime(newScheduleTimeWheel))
	err := s.backend.UpdateTimeWheel(newScheduleTimeWheel)
	if err != nil {
		log.Error().Msgf("❌ error update time wheel for db: %s ", err)
		return err
	}

	s.ScheduleTimeWheel = s.ScheduleTimeWheel.Add(types.TimeWheelSize)
	return nil
}

// UpdateStakeFinalizedStatus defines the periodic process of reading stake records from the database,
// and updating the status of records in the database based on the status of transactions on the BTC chain.
func (s *Server) UpdateStakeFinalizedStatus() {
	timer := time.NewTicker(5 * time.Minute)

	for range timer.C {
		stakeVerifyParams, err := s.backend.QueryNoFinalizedStakeTx()
		if err != nil {
			log.Error().Msgf("💥 error when query no finalized stake tx %s", err)
			continue
		}

		newStatuses, err := s.btcClient.UpdateStakeRecordFinalizedStatus(stakeVerifyParams)
		if err != nil {
			log.Error().Msgf("💥 error when check state records %s", err)
			continue
		}

		for i, status := range newStatuses {
			if status == stakeVerifyParams[i].FinalizedStatus {
				continue
			}

			err := s.backend.UpdateStakeFinalizedStatus(stakeVerifyParams[i].StakerPublicKey, stakeVerifyParams[i].TxID, int(status))
			if err != nil {
				log.Error().Msgf("💥 error when update state finalize status %s", err)
			}
		}
	}
}
