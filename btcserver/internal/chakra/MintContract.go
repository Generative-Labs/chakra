package chakra

import (
	"context"
	"errors"
	"math/big"

	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/account"
	starknetrpc "github.com/NethermindEth/starknet.go/rpc"
	"github.com/NethermindEth/starknet.go/utils"
	"github.com/ethereum/go-ethereum/rpc"
)

func NewChakraProvider(ctx context.Context, rpcURL string) (*starknetrpc.Provider, error) {
	c, err := rpc.DialContext(ctx, rpcURL)
	if err != nil {
		return nil, err
	}
	return starknetrpc.NewProvider(c), nil
}

func NewChakraAccount(privateKey, publicKey, accountAddr string, provider *starknetrpc.Provider) (*account.Account, error) {
	// Here we are converting the account address to felt
	accountAddress, err := utils.HexToFelt(accountAddr)
	if err != nil {
		return nil, err
	}

	// Initializing the account memkeyStore
	ks := account.NewMemKeystore()
	fakePrivKeyBI, ok := new(big.Int).SetString(privateKey, 0)
	if !ok {
		return nil, errors.New("invalid Private Key")
	}

	ks.Put(publicKey, fakePrivKeyBI)

	account, err := account.NewAccount(provider, accountAddress, publicKey, ks, 0)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func RPCCall(provider *starknetrpc.Provider, contractAddressHex string, method string) ([]*felt.Felt, error) {
	contractAddress, err := utils.HexToFelt(contractAddressHex)
	if err != nil {
		panic(err)
	}

	// Make read contract call
	tx := starknetrpc.FunctionCall{
		ContractAddress:    contractAddress,
		EntryPointSelector: utils.GetSelectorFromNameFelt(method),
	}

	callResp, err := provider.Call(context.Background(), tx, starknetrpc.BlockID{Tag: "latest"})
	if err != nil {
		return nil, err
	}
	return callResp, nil
}

func RewardTo(ctx context.Context, cAcctount *account.Account, contractAddressHex string, amount string) (*starknetrpc.AddInvokeTransactionResponse, error) {
	contractAddress, err := utils.HexToFelt(contractAddressHex)
	if err != nil {
		panic(err)
	}

	// Getting the nonce from the account
	nonce, err := cAcctount.Nonce(ctx, starknetrpc.BlockID{Tag: "latest"}, cAcctount.AccountAddress)
	if err != nil {
		return nil, err
	}

	// Here we are setting the maxFee
	maxfee, err := utils.HexToFelt("0x574fbde6000")
	if err != nil {
		return nil, err
	}

	// Building the InvokeTx struct
	invokeTx := starknetrpc.InvokeTxnV1{
		MaxFee:        maxfee,
		Version:       starknetrpc.TransactionV1,
		Nonce:         nonce,
		Type:          starknetrpc.TransactionType_Invoke,
		SenderAddress: cAcctount.AccountAddress,
	}
	callData := make([]*felt.Felt, 0)
	callData = append(callData, AmountToFelt(amount)...)

	// Make read contract call
	fnCall := starknetrpc.FunctionCall{
		ContractAddress:    contractAddress,
		EntryPointSelector: utils.GetSelectorFromNameFelt("RewardTo"),
		Calldata:           callData,
	}

	// Building the Calldata with the help of FmtCalldata where we pass in the FnCall struct along with the Cairo version
	invokeTx.Calldata, err = cAcctount.FmtCalldata([]starknetrpc.FunctionCall{fnCall})
	if err != nil {
		return nil, err
	}

	// Signing of the transaction that is done by the account
	err = cAcctount.SignInvokeTransaction(ctx, &invokeTx)
	if err != nil {
		return nil, err
	}

	// After the signing we finally call the AddInvokeTransaction in order to invoke the contract function
	resp, err := cAcctount.AddInvokeTransaction(ctx, invokeTx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SubmitTXInfo tx_id, btc_amount, expire_at, recipient_address
func SubmitTXInfo(ctx context.Context, cAcctount *account.Account, contractAddressHex string, txID string, amount string, expireAt uint64, rewardReceiver string) (*starknetrpc.AddInvokeTransactionResponse, error) {
	contractAddress, err := utils.HexToFelt(contractAddressHex)
	if err != nil {
		panic(err)
	}

	// Getting the nonce from the account
	nonce, err := cAcctount.Nonce(ctx, starknetrpc.BlockID{Tag: "latest"}, cAcctount.AccountAddress)
	if err != nil {
		return nil, err
	}

	// Here we are setting the maxFee
	maxfee, err := utils.HexToFelt("0x574fbde6000")
	if err != nil {
		return nil, err
	}

	// Building the InvokeTx struct
	invokeTx := starknetrpc.InvokeTxnV1{
		MaxFee:        maxfee,
		Version:       starknetrpc.TransactionV1,
		Nonce:         nonce,
		Type:          starknetrpc.TransactionType_Invoke,
		SenderAddress: cAcctount.AccountAddress,
	}

	callData := make([]*felt.Felt, 0)
	callData = append(callData, AmountToFelt(txID)...)
	callData = append(callData, AmountToFelt(amount)...)
	callData = append(callData, AmountToFelt(big.NewInt(int64(expireAt)).String())...)
	callData = append(callData, AddressToFelt(rewardReceiver))

	// Make read contract call
	fnCall := starknetrpc.FunctionCall{
		ContractAddress:    contractAddress,
		EntryPointSelector: utils.GetSelectorFromNameFelt("SubmitTXInfo"),
		Calldata:           callData,
	}

	// Building the Calldata with the help of FmtCalldata where we pass in the FnCall struct along with the Cairo version
	invokeTx.Calldata, err = cAcctount.FmtCalldata([]starknetrpc.FunctionCall{fnCall})
	if err != nil {
		return nil, err
	}

	// Signing of the transaction that is done by the account
	err = cAcctount.SignInvokeTransaction(ctx, &invokeTx)
	if err != nil {
		return nil, err
	}

	// After the signing we finally call the AddInvokeTransaction in order to invoke the contract function
	resp, err := cAcctount.AddInvokeTransaction(ctx, invokeTx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func Uint256ToFelt252(x *big.Int) (*big.Int, *big.Int) {
	const FeltSize = 252

	// Calculate the maximum value of a Felt
	maxFelt := new(big.Int).Exp(big.NewInt(2), big.NewInt(FeltSize), nil)

	// Perform modulo operation to ensure the value fits within Felt size
	firstPart := new(big.Int).Mod(x, maxFelt)
	secondPart := new(big.Int).Rsh(x, FeltSize)

	return firstPart, secondPart
}

func AddressToFelt(addr string) *felt.Felt {
	return utils.BigIntToFelt(utils.StrToBig(addr))
}

func AmountToFelt(amount string) []*felt.Felt {
	firstPart, secondPart := Uint256ToFelt252(utils.StrToBig(amount))
	return []*felt.Felt{utils.BigIntToFelt(firstPart), utils.BigIntToFelt(secondPart)}
}
