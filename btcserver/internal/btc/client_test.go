package btc_test

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

const (
	bobPrivateKey   = "5JoQtsKQuH8hC9MyvfJAqo6qmKLm8ePYNucs7tPu2YxG12trzBt"
	alicePrivateKey = "5KQr5T79wjyCzaWjMCw3wXB9VF93Wvuk8FkEb36siVXdChmW8v7"
)

func Test_RunTimeLock(_ *testing.T) {
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

	amount, err := btcutil.NewAmount(0.001)
	if err != nil {
		fmt.Println(err)
		return
	}
	hash, err := rpcClient.SendToAddress(address, amount)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hash)

	// Get transaction from transaction hash
	tx, err := rpcClient.GetRawTransaction(hash)
	if err != nil {
		fmt.Println(err)
		return
	}

	prevOutIndex := uint32(1)
	csvMsgTx := wire.NewMsgTx(2)
	prevOut := tx.MsgTx().TxOut[prevOutIndex].PkScript
	txIn := wire.NewTxIn(&wire.OutPoint{Hash: *hash, Index: prevOutIndex}, nil, nil)

	csvMsgTx.AddTxIn(txIn)

	builder := txscript.NewScriptBuilder()
	builder.AddInt64(20)
	builder.AddOp(txscript.OP_CHECKSEQUENCEVERIFY)
	builder.AddOp(txscript.OP_DROP)
	builder.AddData(bobPubKey.SerializeCompressed())
	builder.AddOp(txscript.OP_CHECKSIG)

	csvScript, err := builder.Script()
	if err != nil {
		fmt.Println(err)
		return
	}

	h := sha256.Sum256(csvScript)
	fmt.Println(hex.EncodeToString(h[:]))

	p2shCSVAddr, err := btcutil.NewAddressScriptHash(csvScript, &chaincfg.RegressionNetParams)
	if err != nil {
		fmt.Println(err)
		return
	}

	p2CSVshScript, err := txscript.PayToAddrScript(p2shCSVAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	txOut := wire.NewTxOut(int64(50000), p2CSVshScript)
	csvMsgTx.AddTxOut(txOut)

	sigScript, err := txscript.SignatureScript(csvMsgTx, 0, prevOut, txscript.SigHashAll, bobPrivKey.PrivKey, true)
	if err != nil {
		fmt.Println(err)
		return
	}

	csvMsgTx.TxIn[0].SignatureScript = sigScript

	if err := validateTransaction(csvMsgTx, 0, prevOut); err != nil {
		fmt.Println("Transaction validation failed:", err)
		return
	}

	fmt.Println("Transaction is valid")

	encodeTx, err := encodeTransaction(csvMsgTx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hex.EncodeToString(encodeTx))

	csvHash, err := rpcClient.SendRawTransaction(csvMsgTx, false)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("csv tx hash  %s\n", csvHash.String())

	spendCsvMsgTx := wire.NewMsgTx(2)
	csvTx, err := rpcClient.GetRawTransaction(csvHash)
	if err != nil {
		fmt.Println(err)
		return
	}
	csvPrevOut := csvTx.MsgTx().TxOut[0].PkScript
	selfAddrScript, err := txscript.PayToAddrScript(address)
	if err != nil {
		fmt.Println(err)
		return
	}
	csvTxIn := wire.NewTxIn(&wire.OutPoint{Hash: *csvHash, Index: 0}, nil, nil)
	csvTxIn.Sequence = 30
	spendCsvMsgTx.AddTxIn(csvTxIn)
	spendCsvMsgTx.AddTxOut(wire.NewTxOut(int64(10000), selfAddrScript))

	sig, err := txscript.RawTxInSignature(spendCsvMsgTx, 0, csvScript, txscript.SigHashAll, bobPrivKey.PrivKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	signature := txscript.NewScriptBuilder()
	signature.AddData(sig)
	signature.AddData(csvScript)
	redeemSigScript, err := signature.Script()
	if err != nil {
		fmt.Println(err)
		return
	}

	spendCsvMsgTx.TxIn[0].SignatureScript = redeemSigScript

	flags := txscript.ScriptBip16 | txscript.ScriptVerifyDERSignatures |
		txscript.ScriptStrictMultiSig |
		txscript.ScriptDiscourageUpgradableNops | txscript.ScriptVerifyCheckSequenceVerify
	vm, err := txscript.NewEngine(csvPrevOut, spendCsvMsgTx, 0, flags, nil, nil, -1, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := vm.Execute(); err != nil {
		fmt.Println(err)
		return
	}

	ent, err := encodeTransaction(spendCsvMsgTx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(hex.EncodeToString(ent))
	fmt.Println("end")
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
