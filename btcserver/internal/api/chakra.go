package api

import (
	"github.com/generativelabs/btcserver/internal/chakra"
	"github.com/rs/zerolog/log"
	"time"
)

const (
	ReleaseRewardTime = 8
)

var ScheduleSignalChannel = make(chan struct{}, 0)
var ReleaseRewardChannel = make(chan string, 1000)

func (s *Server) ScheduleTask() {
	go func() {
		for {
			now := time.Now()
			if now.Hour() == ReleaseRewardTime && now.Minute() == 0 && now.Second() == 0 {
				log.Debug().Msgf("It's %s o'clock now, and rewards are starting to be released", now.String())
				ScheduleSignalChannel <- struct{}{}
			} else {
				time.Sleep(time.Second)
				continue
			}
		}
	}()

	for {
		select {
		case <-ScheduleSignalChannel:
			allTxs := make([]string, 0)
			for i := 0; i < 1000; i++ {
				txs, err := s.backend.QueryAllNotYetLockedUpTx()
				if err != nil {
					log.Error().Msgf("❌ error query all not release tx: %s ", err)
					if i%10 == 0 {
						time.Sleep(time.Second * 5)
					}
					continue
				}
				allTxs = txs
			}

			for _, tx := range allTxs {
				log.Debug().Msgf("Txid %s enters the queue for releasing rewards", tx)
				ReleaseRewardChannel <- tx
			}
		}
	}
}

func (s *Server) AddressesMint() {
	for {
		select {
		case tx := <-ReleaseRewardChannel:
			log.Debug().Msgf("Txid %s starts to initiate mint", tx)

			//mint
			res, err := chakra.ChakraRewardTo(s.Ctx, s.ChakraAccount, s.ContractAddress, tx)
			if err != nil {
				log.Error().Msgf("❌ error mint call: %s ", err)
				continue
			}

			log.Info().Msgf("Chakra reward to success, tx hash: %s ", res.TransactionHash)
		}
	}
}
