package chakra

import (
	"context"

	starknetrpc "github.com/NethermindEth/starknet.go/rpc"
	"github.com/btcsuite/btcd/wire"
	"github.com/ethereum/go-ethereum/rpc"
)

func NewChakraProvider(ctx context.Context, rpcURL string) (*starknetrpc.Provider, error) {
	c, err := rpc.DialContext(ctx, rpcURL)
	if err != nil {
		return nil, err
	}

	p := starknetrpc.NewProvider(c)
	if err != nil {
		return nil, err
	}
	return p, nil
}

type Client struct {
	provider starknetrpc.Provider
}

func NewClient(cfg *RPCConfig) (*Client, error) {
	provider, err := NewChakraProvider(context.Background(), cfg.URL)
	if err != nil {
		return nil, err
	}
	c := Client{
		provider: *provider,
	}
	return &c, nil
}

// LatestSyncedBtcBlockHeight defines a method to retrieve the latest synced BTC
// block header from the BTC header queue contract on the Chakra network.
func (c Client) LatestSyncedBtcBlockHeader() (int64, error) {
	// TODO: implement
	return 0, nil
}

func (c Client) SyncBtcBlockHeader(_ *wire.BlockHeader) error {
	return nil
}
