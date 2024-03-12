package api

import (
	"github.com/generativelabs/btcserver/internal/types"
	"strconv"
	"sync"
	"time"

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
		txs, err := s.backend.QueryAllNotYetLockedUpTxNextFourHours(s.ScheduleTimeWheel.UnixMilli())
		if err != nil {
			log.Error().Msgf("❌ error query all not release tx: %s ", err)
			time.Sleep(time.Second)
			continue
		}

		go s.RewardTasksSchedule(txs)
	}
}

func (s *Server) RewardTasksSchedule(txs []*types.ReleaseTxsInfo) {
	wg := sync.WaitGroup{}
	wg.Add(len(txs))
	for _, tx := range txs {
		go func(tx *types.ReleaseTxsInfo) {
			defer wg.Done()
			err := s.RewardTasks(tx)
			if err != nil {
				log.Error().Msgf("❌ error reward tasks: %s ", err)
				return
			}
		}(tx)
	}

	wg.Wait()

	newScheduleTimeWheel := s.ScheduleTimeWheel.UnixMilli() + 5*time.Minute.Milliseconds()
	err := s.backend.UpdateTimeWheel(newScheduleTimeWheel)
	if err != nil {
		log.Error().Msgf("❌ error update time wheel for db: %s ", err)
		return
	}
	s.ScheduleTimeWheel = s.ScheduleTimeWheel.Add(5 * time.Minute)
}

func (s *Server) RewardTasks(tx *types.ReleaseTxsInfo) error {
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
