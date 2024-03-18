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

		v, err := strconv.Atoi(tw.Value)
		if err != nil {
			log.Error().Msgf("❌ error strconv.Atoi time wheel: %s ", err)
			return
		}

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

	log.Info().Msgf("🔵 Start time wheel schedule, time wheel: %s", s.ScheduleTimeWheel)

	for {
		// Start from the global timestamp, query every 5 minutes, and query all users to be released within 5 minutes
		txs, err := s.backend.QueryAllNotYetLockedUpTxNextPeriod(s.ScheduleTimeWheel.UnixNano(), types.TimeWheelSize)
		if err != nil {
			log.Error().Msgf("❌ error query all not release tx: %s ", err)
			time.Sleep(time.Second)
			continue
		}

		if len(txs) == 0 {
			if time.Since(s.ScheduleTimeWheel) > types.TimeWheelSize {
				err = s.UpdateTimeWheelForDB()
				if err != nil {
					log.Error().Msgf("❌ error update time wheel: %s ", err)
				} else {
					s.UpdateTimeWheel()
				}

				continue
			}

			oldScheduleTimeWheel := s.ScheduleTimeWheel
			log.Info().Msgf("🔵 No tx to be released, now %s sleep Until next Wheel %s", time.Now().String(), oldScheduleTimeWheel.Add(types.TimeWheelSize).String())
			time.Sleep(time.Until(oldScheduleTimeWheel.Add(types.TimeWheelSize)))
			err = s.UpdateTimeWheelForDB()
			if err != nil {
				log.Error().Msgf("❌ error update time wheel: %s ", err)
			} else {
				s.UpdateTimeWheel()
			}
			continue
		}

		log.Info().Msgf("🔵 In the next time period %s - %s, find %d txids that are about to be released: %v", s.ScheduleTimeWheel, s.ScheduleTimeWheel.Add(types.TimeWheelSize), len(txs), txs)
		s.RewardTasksSchedule(txs)
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

	_ = s.UpdateTimeWheelForDB()
	s.UpdateTimeWheel()
}

func (s *Server) RewardTasks(txs []*types.ReleaseTxsInfo) error {
	rt := utils.TimestampToTime(txs[0].ReleasingTime)
	timer := time.NewTimer(time.Until(rt))
	log.Info().Msgf("🔵 Start the timer and prepare to wait %v min [%v]", time.Until(rt).Minutes(), txs[0])

	for { //nolint
		select {
		case <-timer.C:
			txIDs := make([]string, 0)
			for _, tx := range txs {
				txIDs = append(txIDs, tx.Tx)
			}

			log.Info().Msgf("🔵 Start reward to %v ", txIDs)
			res, err := chakra.RewardTo(s.Ctx, s.ChakraAccount, s.ContractAddress, txIDs)
			if err != nil {
				log.Error().Msgf("❌ error reward to txIDs %v: %s ", txIDs, err)
				// todo deal err task
				return err
			}
			log.Info().Msgf("🔵 Chakra reward success, txs hash: %s ", res.TransactionHash)

			for _, tx := range txs {
				log.Info().Msgf("🔵 Start update stake ReleasingTime %v ", tx)
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

func (s *Server) UpdateTimeWheel() {
	newScheduleTimeWheel := utils.TimeTOTimestamp(s.ScheduleTimeWheel) + types.TimeWheelSize.Nanoseconds()
	log.Info().Msgf("Update time wheel to %s ", utils.TimestampToTime(newScheduleTimeWheel))

	s.ScheduleTimeWheel = s.ScheduleTimeWheel.Add(types.TimeWheelSize)
}

func (s *Server) UpdateTimeWheelForDB() error {
	newScheduleTimeWheel := utils.TimeTOTimestamp(s.ScheduleTimeWheel) + types.TimeWheelSize.Nanoseconds()
	log.Info().Msgf("DB update time wheel to %s ", utils.TimestampToTime(newScheduleTimeWheel))
	err := s.backend.UpdateTimeWheel(newScheduleTimeWheel)
	if err != nil {
		log.Error().Msgf("❌ error update time wheel for db: %s ", err)
		return err
	}
	return nil
}

// UpdateStakeFinalizedStatus defines the periodic process of reading stake records from the database,
// and updating the status of records in the database based on the status of transactions on the BTC chain.
func (s *Server) UpdateStakeFinalizedStatus() {
	timer := time.NewTicker(5 * time.Minute)

	log.Info().Msgf("🔨 Start update stake finalized status with a 5 minute ticker")

	for range timer.C {
		oldStateRecords, err := s.backend.QueryNoFinalizedStakeTx()
		if err != nil {
			log.Error().Msgf("💥 error when query no finalized stake tx %s", err)
			continue
		}

		if len(oldStateRecords) == 0 {
			log.Info().Msgf("💤 there is't no finalize states from db.")
			continue
		}
		log.Info().Msgf("🔨get %d no finalize states from db. prepare check them", len(oldStateRecords))

		newStateRecords, err := s.btcClient.UpdateStakeRecords(oldStateRecords)
		if err != nil {
			log.Error().Msgf("💥 error when check state records %s", err)
			continue
		}

		for i, record := range newStateRecords {
			staker := oldStateRecords[i].Staker
			txID := oldStateRecords[i].TxID
			amount := strconv.Itoa(int(oldStateRecords[i].Amount))
			start := record.Start
			deadline := record.Start + oldStateRecords[i].Duration
			releasingTime := record.Start + 24*time.Hour.Nanoseconds()
			rewardReceiver := oldStateRecords[i].RewardReceiver

			if record.Status == oldStateRecords[i].FinalizedStatus {
				continue
			}

			err := s.backend.UpdateStakeFinalizedStatus(staker, txID, int(record.Status), record.Start, deadline, releasingTime)
			if err != nil {
				log.Error().Msgf("💥 error when update state finalize status %s", err)
				continue
			}

			if record.Status == types.TxFinalized {
				log.Info().Msgf("🔵 start submit btc stake tx info to chakra: txID %s amount %s start %d deadline %d rewardReceiver %s",
					txID, amount, start, deadline, rewardReceiver)
				res, err := chakra.SubmitTXInfo(s.Ctx, s.ChakraAccount, s.ContractAddress, txID, amount, start, deadline, rewardReceiver)
				if err != nil {
					log.Error().Msgf("💥 error submit stake tx infos %s", err)
					continue
				}
				err = s.backend.UpdateCanBeSubmitStatus(staker, txID, 1)
				if err != nil {
					log.Error().Msgf("💥 error update stake %s txid %s txhash %s submit status %s",
						staker, txID, res.TransactionHash, err)
				}
				log.Info().Msgf("🔵 Submit btc stake tx info to chakra successfully: txhash %s txID %s",
					res.TransactionHash, txID)
			}
		}
	}
}
