package btc_test

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/generativelabs/btcserver/internal/btc"
	"github.com/generativelabs/btcserver/internal/types"
)

const (
	bobPrivateKey = "5JoQtsKQuH8hC9MyvfJAqo6qmKLm8ePYNucs7tPu2YxG12trzBt"
)

// Example_btc_lock_and_redeem tests that build and send a csv tx, then use a
// redeem tx to spend it.
// You need:
// 1. launch the bitcoin regtest network locally.
// 2. create a btc wallet.
// 3. mint some btc block, then the wallet can get some test bitcoin.
func Example_btc_lock_and_redeem() {
	rpcClient, err := rpcclient.New(&rpcclient.ConnConfig{
		Host:         "localhost:18332",
		User:         "rpcuser",
		Pass:         "rpcpass",
		HTTPPostMode: true,
		DisableTLS:   true,
	}, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	bobPrivKey, err := btcutil.DecodeWIF(bobPrivateKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	bobPubKey := bobPrivKey.PrivKey.PubKey()

	address, err := btcutil.NewAddressPubKey(bobPubKey.SerializeCompressed(), &chaincfg.RegressionNetParams)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(address)

	amount, err := btcutil.NewAmount(0.001)
	if err != nil {
		fmt.Println(err)
		return
	}
	hash, err := rpcClient.SendToAddress(address, amount)
	if err != nil {
		fmt.Println(err)
	}

	// Get transaction from transaction hash
	tx, err := rpcClient.GetRawTransaction(hash)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(tx)

	txRes, err := rpcClient.GetRawTransactionVerbose(hash)
	if err != nil {
		fmt.Println(err)
		return
	}
	jtx1, _ := json.Marshal(txRes)
	fmt.Println(string(jtx1))

	csvMsgTx, csvScript, err := createCSVLockTx(tx.MsgTx(), *hash, 1, 20, 50000, bobPrivKey.PrivKey, bobPubKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	csvHash, err := rpcClient.SendRawTransaction(csvMsgTx, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	txRes2, err := rpcClient.GetRawTransactionVerbose(csvHash)
	if err != nil {
		fmt.Println(err)
		return
	}
	jtx2, _ := json.Marshal(txRes2)
	fmt.Println("////")
	fmt.Println(string(jtx2))

	spendCsvMsgTx, err := createCSVRedeemTx(*csvHash, address, csvScript, bobPrivKey.PrivKey, 10000)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := validateTransaction(spendCsvMsgTx, 0, csvMsgTx.TxOut[0].PkScript); err != nil {
		fmt.Println(err)
		return
	}

	spendTx, err := encodeTransaction(spendCsvMsgTx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("end")
	fmt.Println(hex.EncodeToString(spendTx))

	// Output:
	// end
}

func validateTransaction(msg *wire.MsgTx, idx int, prevOutScript []byte) error {
	vm, err := txscript.NewEngine(prevOutScript, msg, idx, txscript.StandardVerifyFlags, nil, nil, -1, nil)
	if err != nil {
		return err
	}

	if err := vm.Execute(); err != nil {
		return err
	}

	return nil
}

func encodeTransaction(tx *wire.MsgTx) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := tx.Serialize(buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func createCSVLockTx(inTx *wire.MsgTx, inTxID chainhash.Hash, prevOutIndex uint32,
	lockSequence int64, lockAmount int64, senderPrivateKey *secp256k1.PrivateKey,
	receiverPublicKey *secp256k1.PublicKey,
) (*wire.MsgTx, []byte, error) {
	prevOut := inTx.TxOut[prevOutIndex].PkScript

	csvMsgTx := wire.NewMsgTx(2)
	txIn := wire.NewTxIn(&wire.OutPoint{Hash: inTxID, Index: prevOutIndex}, nil, nil)
	csvMsgTx.AddTxIn(txIn)

	builder := txscript.NewScriptBuilder()
	builder.AddInt64(lockSequence)
	builder.AddOp(txscript.OP_CHECKSEQUENCEVERIFY)
	builder.AddOp(txscript.OP_DROP)
	builder.AddData(receiverPublicKey.SerializeCompressed())
	builder.AddOp(txscript.OP_CHECKSIG)

	csvScript, err := builder.Script()
	if err != nil {
		return nil, nil, err
	}

	p2shCSVAddr, err := btcutil.NewAddressScriptHash(csvScript, &chaincfg.RegressionNetParams)
	if err != nil {
		return nil, nil, err
	}

	p2CSVshScript, err := txscript.PayToAddrScript(p2shCSVAddr)
	if err != nil {
		return nil, nil, err
	}

	txOut := wire.NewTxOut(lockAmount, p2CSVshScript)
	csvMsgTx.AddTxOut(txOut)

	sigScript, err := txscript.SignatureScript(csvMsgTx, 0, prevOut, txscript.SigHashAll, senderPrivateKey, true)
	if err != nil {
		return nil, nil, err
	}

	csvMsgTx.TxIn[0].SignatureScript = sigScript

	if err := validateTransaction(csvMsgTx, 0, prevOut); err != nil {
		return nil, nil, err
	}

	return csvMsgTx, csvScript, nil
}

func createCSVRedeemTx(lockTxHash chainhash.Hash, receiverAddress btcutil.Address,
	lockScript []byte, signPrivKey *secp256k1.PrivateKey, redeemAmount int64,
) (*wire.MsgTx, error) {
	spendCsvMsgTx := wire.NewMsgTx(2)
	outP2PKScript, err := txscript.PayToAddrScript(receiverAddress)
	if err != nil {
		return nil, err
	}

	csvTxIn := wire.NewTxIn(&wire.OutPoint{Hash: lockTxHash, Index: 0}, nil, nil)
	csvTxIn.Sequence = 30
	spendCsvMsgTx.AddTxIn(csvTxIn)
	spendCsvMsgTx.AddTxOut(wire.NewTxOut(redeemAmount, outP2PKScript))

	sig, err := txscript.RawTxInSignature(spendCsvMsgTx, 0, lockScript, txscript.SigHashAll, signPrivKey)
	if err != nil {
		return nil, err
	}
	signature := txscript.NewScriptBuilder()
	signature.AddData(sig)
	signature.AddData(lockScript)
	redeemSigScript, err := signature.Script()
	if err != nil {
		return nil, err
	}

	spendCsvMsgTx.TxIn[0].SignatureScript = redeemSigScript

	return spendCsvMsgTx, nil
}

// Example_checkStakeTxs tests check stake txs.
func Example_checkStakeTxs() {
	client, err := btc.NewClient(btc.Config{
		NetworkName: chaincfg.TestNet3Params.Name,
		RPCHost:     "bitcoin-testnet.blastapi.io/97d216a1-c44b-461d-9624-2231e517a4c6",
		RPCUser:     "rpcuser",
		RPCPass:     "rpcpass",
		DisableTLS:  false,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	stakeRecord := types.StakeVerificationParam{
		TxID:            "9750a37c3713cef0eb5169c932397589b82f15abc1af7354e62d1cf32263b708",
		StakerPublicKey: "03d87175c1ca3222d1500def9e79692fbd924b85c83e784907b1d1babded7cc72e",
		Amount:          1000,
		Duration:        7 * (24 * time.Hour.Nanoseconds()),
		FinalizedStatus: types.TxPending,
	}

	res, err := client.UpdateStakeRecords([]*types.StakeVerificationParam{&stakeRecord})
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(res) != 1 {
		fmt.Println(err)
		return
	}

	if res[0].Status != types.Mismatch {
		fmt.Println("check include tx failed")
		return
	}

	fmt.Println("end")

	// Output:
	// end
}

const (
	messageSignatureHeader = "Bitcoin Signed Message:\n"
)

func Example_signAndVerifyMessage() {
	bobPrivKey, err := btcutil.DecodeWIF(bobPrivateKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	bobPk := "024edfcf9dfe6c0b5c83d1ab3f78d1b39a46ebac6798e08e19761f5ed89ec83c10"

	message := "0x65fbbc6ed72f28f38e9b7b440b4115b143a35cfe7ceb390f448fa0a1bcbd8dc/1710587088607000"
	var buf bytes.Buffer
	_ = wire.WriteVarString(&buf, 0, messageSignatureHeader)
	_ = wire.WriteVarString(&buf, 0, message)
	messageHash := chainhash.DoubleHashB(buf.Bytes())

	sigB, err := ecdsa.SignCompact(bobPrivKey.PrivKey, messageHash, true)
	if err != nil {
		fmt.Println(err)
		return
	}

	recoverPK, ok, err := ecdsa.RecoverCompact(sigB, messageHash)
	if !ok {
		fmt.Println("no compact signature")
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}

	addressPK, err := btcutil.NewAddressPubKey(recoverPK.SerializeCompressed(), &chaincfg.TestNet3Params)
	if err != nil {
		fmt.Println(err)
		return
	}

	if addressPK.String() != strings.TrimPrefix(bobPk, "0x") {
		fmt.Println("mismatch public key")
		return
	}

	fmt.Println("end")
	// Output:
	// end
}
