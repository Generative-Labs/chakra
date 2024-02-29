package headeroracle

import (
	"time"

	"github.com/btcsuite/btcd/wire"
	"github.com/generative-labs/relayer/internal/btc"
	"github.com/generative-labs/relayer/internal/chakra"
	"github.com/rs/zerolog/log"
)

const (
	newBlockWaitTimeout = 60 * time.Second
	blockHeaderChanLen  = 1024
)

type HeaderOracle struct {
	btcClient    *btc.Client
	chakraClient *chakra.Client
	oracleConfig OracleConfig

	blockHeaderCh chan *wire.BlockHeader
	stopCh        chan struct{}
}

func New(cfg *Config) (*HeaderOracle, error) {
	btcClient, err := btc.NewClient(&cfg.BtcRPC)
	if err != nil {
		return nil, err
	}

	chakraClient, err := chakra.NewClient(&cfg.Chakra.RPC)
	if err != nil {
		return nil, err
	}

	headerOracle := &HeaderOracle{
		btcClient:     btcClient,
		chakraClient:  chakraClient,
		oracleConfig:  cfg.Oracle,
		blockHeaderCh: make(chan *wire.BlockHeader, blockHeaderChanLen),
		stopCh:        make(chan struct{}),
	}

	return headerOracle, nil
}

func (h HeaderOracle) Start() error {
	errCh := make(chan error)

	go func() {
		errCh <- h.retrieveBtcBlockHeader()
	}()

	go func() {
		errCh <- h.syncBtcBlockHeaderToChakra()
	}()

	select {
	case err := <-errCh:
		return err
	case <-h.stopCh:
		log.Info().Msg("📴 HeaderOracle receive stop signal")
		return nil
	}
}

func (h HeaderOracle) retrieveBtcBlockHeader() error {
	lastSyncedHeight, err := h.chakraClient.LatestSyncedBtcBlockHeader()
	if err != nil {
		return err
	}

	for {
		curBtcHeight, err := h.btcClient.LatestBlock()
		if err != nil {
			return err
		}

		if curBtcHeight == lastSyncedHeight {
			log.Info().Msg("👏 already sync to latest btc block")
			time.Sleep(newBlockWaitTimeout)
			continue
		}

		if curBtcHeight < lastSyncedHeight {
			log.Fatal().Msg("💥 The BTC block height on Chakra exceeds that obtained" +
				"from the current BTC RPC. The RPC may be experiencing issues.")
		}

		for i := lastSyncedHeight + 1; i < curBtcHeight; i++ {
			header, err := h.btcClient.GetBlockHeader(i)
			if err != nil {
				log.Error().Msgf("❌ Get block header from btc rpc failed, error: %s", err)
				continue
			}

			h.blockHeaderCh <- header
			lastSyncedHeight = i
		}

		time.Sleep(time.Duration(h.oracleConfig.PollInterval))
	}
}

func (h HeaderOracle) syncBtcBlockHeaderToChakra() error {
	for {
		header := <-h.blockHeaderCh
		if err := h.chakraClient.SyncBtcBlockHeader(header); err != nil {
			return err
		}
	}
}

func (h HeaderOracle) Stop() {
	h.stopCh <- struct{}{}
}
