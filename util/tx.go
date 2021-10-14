package util
import (
	"encoding/json"
	"fmt"
	"github.com/33cn/chain33/common"
	"github.com/33cn/chain33/rpc/jsonclient"
	"github.com/33cn/chain33/types"
	rpctypes "github.com/33cn/chain33/rpc/types"
	cty "github.com/33cn/chain33/system/dapp/coins/types"
	pty "github.com/33cn/plugin/plugin/dapp/token/types"
	"github.com/pkg/errors"
	"github.com/golang/protobuf/proto"
)


func DecodeHexTransaction(hexTx string) (*types.Transaction, error) {
	param := types.ReqDecodeRawTransaction{TxHex: hexTx}
	return DecodeRawTransaction(param)
}


func DecodeRawTransaction(param types.ReqDecodeRawTransaction) (*types.Transaction, error) {
	var tx types.Transaction

	txBytes, err := common.FromHex(param.TxHex)
	if err != nil {
		return nil, errors.Wrap(err, "common.FromHex")
	}
	err = types.Decode(txBytes, &tx)
	if err != nil {
		return nil, errors.Wrap(err, "Decode txBytes")
	}
	return &tx, nil
}


func ChangeTxToAddress(tx *types.Transaction, newToAddr string) *types.Transaction {
	tx.To = newToAddr
	return tx
}

func ChangeTxAmt(tx *types.Transaction, newAmt float64) (*types.Transaction, error) {
	if string(tx.Execer) == "coins" {
		action := &cty.CoinsAction{}
		err := types.Decode(tx.GetPayload(), action)
		if err != nil {
			return nil, errors.Wrap(err, "decode tx payload")
		}

		transfer := action.GetTransfer()
		transfer.Amount = int64((newAmt + 1e-9) * 1e8)

		v := &cty.CoinsAction_Transfer{transfer}
		//coinTransfer := &types.CoinsAction{Value: v, Ty: types.CoinsActionTransfer}
		action.Value = v
		action.Ty = cty.CoinsActionTransfer
		//tx = &types.Transaction{Execer: tx.Execer, Payload: types.Encode(action),Fee:tx.Fee, To: tx.To, Nonce: tx.Nonce}

		tx.Payload = types.Encode(action)

	} else {
		action := &pty.TokenAction{}
		err := types.Decode(tx.GetPayload(), action)
		if err != nil {
			return nil, errors.Wrap(err, "decode tx payload ")
		}
		transfer := action.GetTransfer()
		transfer.Amount = int64((newAmt + 1e-9) * 1e8)
		v := &pty.TokenAction_Transfer{Transfer: transfer}

		action.Value = v
		action.Ty = pty.ActionTransfer
		tx.Payload = types.Encode(action)

	}
	return tx, nil
}

func ChangeTx(tx *types.Transaction, newToAddr string, newAmt float64) (*types.Transaction, error)  {
	if tx == nil {
		return nil, errors.New("tx is nil")
	}

	tmp := ChangeTxToAddress(tx, newToAddr)
	newTx, err := ChangeTxAmt(tmp, newAmt)
	if err != nil {
		return nil, err
	}
	return newTx, nil
}

func PrintTxDetail(trans *types.Transaction) error {
	txCopy := *trans
	var tx = &txCopy
	res, err := rpctypes.DecodeTx(tx, types.DefaultCoinPrecision)
	if err != nil {
		return errors.Wrap(err, "rpctypes.DecodeTx")
	}
	var payload proto.Message
	if string(tx.Execer )== "coins" {
		payload = &cty.CoinsAction{}
	} else {
		payload = &pty.TokenAction{}
	}


	err = types.Decode(tx.GetPayload(), payload)
	if err != nil {
		return errors.Wrap(err, "types.Decode tx payload")
	}
	var pljson json.RawMessage
	if payload != nil {
		//fmt.Println("ptypes.PBToJSONUTF8 pl", payload )
		pljson, _ = types.PBToJSONUTF8(payload)
	}

	res.Payload = pljson
	txBytes, err := json.MarshalIndent(res, "", "\t")
	fmt.Printf("%s\n", txBytes)
	
	return nil
}

func PrintHexTxDetail(hexTx string) error {
	tx, err := DecodeHexTransaction(hexTx)
	if err != nil {
		fmt.Println(err)
	}
	err = PrintTxDetail(tx)
	if err != nil {
		return err
	}
	return nil
}

func EncodeToString(tx *types.Transaction) (string) {
	return common.ToHex(types.Encode(tx))
}

func SendTransaction(hexTx string, rpc_addr string) (string, error) {
	jsonCli, err := jsonclient.NewJSONClient(rpc_addr)
	if err != nil {
		fmt.Println("NewJSONClient err:", err)
		return "", err
	}

	param := rpctypes.RawParm{
		Data:  hexTx,
	}

	var result string
	err = jsonCli.Call("Chain33.SendTransaction", param, &result)

	return result, err
}
