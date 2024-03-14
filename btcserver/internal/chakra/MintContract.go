package chakra

import (
	"context"
	"errors"
	"math/big"

	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/account"
	"github.com/NethermindEth/starknet.go/curve"
	starknetrpc "github.com/NethermindEth/starknet.go/rpc"
	"github.com/NethermindEth/starknet.go/utils"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/rs/zerolog/log"
)

func NewChakraProvider(ctx context.Context, rpcURL string) (*starknetrpc.Provider, error) {
	c, err := rpc.DialContext(ctx, rpcURL)
	if err != nil {
		return nil, err
	}
	return starknetrpc.NewProvider(c), nil
}

func NewChakraAccount(ctx context.Context, rpcURL string, privateKey, accountAddr string) (*account.Account, error) {
	provider, err := NewChakraProvider(ctx, rpcURL)
	if err != nil {
		log.Fatal().Msgf("❌ Fatal error new chakra provider: %s ", err)
	}

	publicKey := GetPublicKeyFromPrivateKey(privateKey)

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

	acc, err := account.NewAccount(provider, accountAddress, publicKey, ks, 0)
	if err != nil {
		return nil, err
	}

	return acc, nil
}

func RPCCall(cAccount *account.Account, contractAddressHex string, method string, callData []*felt.Felt) ([]*felt.Felt, error) {
	contractAddress, err := utils.HexToFelt(contractAddressHex)
	if err != nil {
		panic(err)
	}

	// Make read contract call
	tx := starknetrpc.FunctionCall{
		ContractAddress:    contractAddress,
		EntryPointSelector: utils.GetSelectorFromNameFelt(method),
		Calldata:           callData,
	}

	callResp, err := cAccount.Call(context.Background(), tx, starknetrpc.BlockID{Tag: "latest"})
	if err != nil {
		return nil, err
	}
	return callResp, nil
}

func RewardTo(ctx context.Context, cAccount *account.Account, contractAddressHex string, txIDs []string) (*starknetrpc.AddInvokeTransactionResponse, error) {
	contractAddress, err := utils.HexToFelt(contractAddressHex)
	if err != nil {
		panic(err)
	}

	// Getting the nonce from the account
	nonce, err := cAccount.Nonce(ctx, starknetrpc.BlockID{Tag: "latest"}, cAccount.AccountAddress)
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
		SenderAddress: cAccount.AccountAddress,
	}

	params := ArrBtcTxIDToFelt(txIDs)

	lenP := utils.BigIntToFelt(big.NewInt(int64(len(txIDs))))
	callData := make([]*felt.Felt, 0)
	callData = append(callData, lenP)
	callData = append(callData, params...)

	// Make read contract call
	fnCall := starknetrpc.FunctionCall{
		ContractAddress:    contractAddress,
		EntryPointSelector: utils.GetSelectorFromNameFelt("RewardTo"),
		Calldata:           callData,
	}

	// Building the Calldata with the help of FmtCalldata where we pass in the FnCall struct along with the Cairo version
	invokeTx.Calldata, err = cAccount.FmtCalldata([]starknetrpc.FunctionCall{fnCall})
	if err != nil {
		return nil, err
	}

	// Signing of the transaction that is done by the account
	err = cAccount.SignInvokeTransaction(ctx, &invokeTx)
	if err != nil {
		return nil, err
	}

	// After the signing we finally call the AddInvokeTransaction in order to invoke the contract function
	resp, err := cAccount.AddInvokeTransaction(ctx, invokeTx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SubmitTXInfo tx_id, btc_amount, startAt, expire_at, recipient_address
func SubmitTXInfo(ctx context.Context, cAccount *account.Account, contractAddressHex string, txID string, amount string, startAt, expireAt uint64, rewardReceiver string) (*starknetrpc.AddInvokeTransactionResponse, error) {
	contractAddress, err := utils.HexToFelt(contractAddressHex)
	if err != nil {
		panic(err)
	}

	// Getting the nonce from the account
	nonce, err := cAccount.Nonce(ctx, starknetrpc.BlockID{Tag: "latest"}, cAccount.AccountAddress)
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
		SenderAddress: cAccount.AccountAddress,
	}

	callData := make([]*felt.Felt, 0)
	callData = append(callData, BtcTxIDToFelt(txID)...)
	callData = append(callData, AmountToFelt(amount)...)
	callData = append(callData, utils.BigIntToFelt(big.NewInt(int64(startAt))))
	callData = append(callData, utils.BigIntToFelt(big.NewInt(int64(expireAt))))
	receiver, err := utils.HexToFelt(rewardReceiver)
	if err != nil {
		return nil, err
	}
	callData = append(callData, receiver)

	// Make read contract call
	fnCall := starknetrpc.FunctionCall{
		ContractAddress:    contractAddress,
		EntryPointSelector: utils.GetSelectorFromNameFelt("SubmitTXInfo"),
		Calldata:           callData,
	}

	// Building the Calldata with the help of FmtCalldata where we pass in the FnCall struct along with the Cairo version
	invokeTx.Calldata, err = cAccount.FmtCalldata([]starknetrpc.FunctionCall{fnCall})
	if err != nil {
		return nil, err
	}

	// Signing of the transaction that is done by the account
	err = cAccount.SignInvokeTransaction(ctx, &invokeTx)
	if err != nil {
		return nil, err
	}

	// After the signing we finally call the AddInvokeTransaction in order to invoke the contract function
	resp, err := cAccount.AddInvokeTransaction(ctx, invokeTx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func GetPublicKeyFromPrivateKey(privateKey string) string {
	privInt := utils.HexToBN(privateKey)

	pubX, _, err := curve.Curve.PrivateToPoint(privInt)
	if err != nil {
		log.Fatal().Msgf("❌ Fatal error generate public key: %s ", err)
		panic(err)
	}

	pubKey := utils.BigToHex(pubX)

	return pubKey
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

func NewBtcTxIDToFelt(txID string) *felt.Felt {
	fb := utils.HexToBN(txID)
	f := utils.BigIntToFelt(fb)
	return f
}

func BtcTxIDToFelt(amount string) []*felt.Felt {
	firstPart, secondPart := Uint256ToFelt252(utils.HexToBN(amount))
	return []*felt.Felt{utils.BigIntToFelt(firstPart), utils.BigIntToFelt(secondPart)}
}

func ArrBtcTxIDToFelt(txID []string) []*felt.Felt {
	newTxIDs := make([]*felt.Felt, 0)

	for _, tx := range txID {
		fb := utils.HexToBN(tx)
		f := utils.BigIntToFelt(fb)
		newTxIDs = append(newTxIDs, f)
	}

	return newTxIDs
}

func AmountToFelt(amount string) []*felt.Felt {
	firstPart, secondPart := Uint256ToFelt252(utils.StrToBig(amount))
	return []*felt.Felt{utils.BigIntToFelt(firstPart), utils.BigIntToFelt(secondPart)}
}
