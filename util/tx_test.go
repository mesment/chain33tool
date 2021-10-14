package util

import (
	"fmt"
	"github.com/33cn/chain33/common"
	"github.com/33cn/chain33/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

var  (
	rpc_addr = "http://127.0.0.1:8801"
	signedTxHex = "0a05636f696e73121718010a130a034254591080efb9341a07746573743132331a6e080112210278ba706d605cbdb0d517e2b38f93711421b29f5d75b9ba4e0ac1aec2ba4a37ce1a473045022100b5ee4acfe9de9d4405aeb1c11a43e9350403b616bbefebd273867be5056a35590220210af2e429100175108269f51da67fa7c6475af01e0e943fdb68616826fb3e4c20b0db0630eb93fbee91dcedd5663a2231505569476362736363667857337a757648585a424a667a6e7a697068356d69416f"
)


func TestChangeTxToAddress(t *testing.T) {

	tx, err := DecodeHexTransaction(signedTxHex)
	assert.NoError(t, err)

	t.Log("---------------Print origin transaction info--------------------------------------")
	PrintTxDetail(tx)

	newToAddr := "16vAAd4WeqWMCJjeTciUTvJHVj9GuPAyWQ"
	t.Logf(fmt.Sprintf("---------------Change transaction to addr to:%s-----------------------\n", newToAddr))
	changedToAddrTx := ChangeTxToAddress(tx, newToAddr)
	t.Log("---------------Print new transaction info------------------------------------------")
	PrintTxDetail(changedToAddrTx)

	hexChangedTx := common.ToHex(types.Encode(changedToAddrTx))
	//t.Log("---------------Print new transaction in hex----------------------------------------------")
	t.Log(hexChangedTx)

	/*
	t.Log("---------------Send new transaction------------------------------------------------")
	result, err := SendTransaction(hexChangedTx, rpc_addr)
	assert.Equal(t, types.ErrSign, err)
	if err != nil {
		t.Log("SendTransaction err:", err)
	} else {
		t.Log("SendTransaction resp:", result)
	}

	 */

}

func TestChangeTxAmt(t *testing.T) {
	tx, err := DecodeHexTransaction(signedTxHex)
	assert.NoError(t, err)

	t.Log("---------------Print origin transaction info-----------------------------------")
	PrintTxDetail(tx)

	newAmt := 2.1
	t.Logf(fmt.Sprintf("---------------Change transaction amount to:%f------------------------\n", newAmt))
	changedAmtTx, err := ChangeTxAmt(tx, newAmt)
	assert.NoError(t, err)
	t.Log("---------------Print new transaction info---------------------------------------")
	PrintTxDetail(changedAmtTx)

	hexChangedTx := common.ToHex(types.Encode(changedAmtTx))
	//t.Log("---------------Print new transaction in hex-------------------------------------")
	t.Log(hexChangedTx)

	/*
	t.Log("---------------Send new transaction---------------------------------------------")

	result, err := SendTransaction(hexChangedTx, rpc_addr)
	assert.Equal(t, types.ErrSign, err)
	if err != nil {
		t.Log("SendTransaction err:", err)
	} else {
		t.Log("SendTransaction resp:", result)
	}

	 */
}