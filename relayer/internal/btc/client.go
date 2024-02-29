package btc

import (
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/wire"
)

func CrateBtcRPCClient(cfg *RPCConfig) (*rpcclient.Client, error) {
	return rpcclient.New(&rpcclient.ConnConfig{
		Host:         cfg.RPCHost,
		User:         cfg.RPCUser,
		Pass:         cfg.RPCPass,
		DisableTLS:   cfg.DisableTLS,
		HTTPPostMode: true,
	}, nil)
}

type Client struct {
	rpcClient *rpcclient.Client
}

func NewClient(cfg *RPCConfig) (*Client, error) {
	rpcClient, err := CrateBtcRPCClient(cfg)
	if err != nil {
		return nil, err
	}
	c := Client{
		rpcClient: rpcClient,
	}
	return &c, nil
}

func (c Client) LatestBlock() (int64, error) {
	return c.rpcClient.GetBlockCount()
}

// GetBlockHeaders defines a method to retrieve block headers from the BTC RPC.
// The interval retrieved is [from, to).
func (c Client) GetBlockHeaders(from, to int64) ([]wire.BlockHeader, error) {
	var blockHeaders []wire.BlockHeader
	for i := from; i < to; i++ {
		hash, err := c.rpcClient.GetBlockHash(i)
		if err != nil {
			return nil, err
		}
		header, err := c.rpcClient.GetBlockHeader(hash)
		if err != nil {
			return nil, err
		}
		blockHeaders = append(blockHeaders, *header)
	}

	return blockHeaders, nil
}

func (c Client) GetBlockHeader(height int64) (*wire.BlockHeader, error) {
	hash, err := c.rpcClient.GetBlockHash(height)
	if err != nil {
		return nil, err
	}
	header, err := c.rpcClient.GetBlockHeader(hash)
	if err != nil {
		return nil, err
	}

	return header, nil
}
