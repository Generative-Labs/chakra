package btc

import "github.com/btcsuite/btcd/rpcclient"

type Client struct {
	rpcClient *rpcclient.Client
}

func New(config Config) (*Client, error) {
	rpcClient, err := rpcclient.New(&rpcclient.ConnConfig{
		Host:         config.RPCHost,
		User:         config.RPCUser,
		Pass:         config.RPCPass,
		DisableTLS:   config.DisableTLS,
		HTTPPostMode: true,
	}, nil)
	if err != nil {
		return nil, err
	}

	client := Client{
		rpcClient: rpcClient,
	}

	return &client, nil
}

func (c *Client) CheckTransactionSignature(rawTx, senderAddress, signature string) {}

func (c *Client) CheckSignature() {}

func (c *Client) CheckTransactionInclusion() {}

func (c *Client) CheckTransactionFinalized() {}
