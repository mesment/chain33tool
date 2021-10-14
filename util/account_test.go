package util

import (
	"github.com/33cn/chain33/types"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

var chain33Cfg *types.Chain33Config
func init()  {
	cfgFile := "./chain33.solo.toml"
	chain33Cfg = types.NewChain33Config(types.MergeCfg(types.ReadFile(cfgFile), ""))
}

func TestAlterBalance(t *testing.T) {
	InitializeDB(chain33Cfg)
	defer CloseDB()
	var addr = "12qyocayNF7Lv6C9qW4avxs2E7U41fKSfv"
	balance, err := GetDBBalance(addr)
	t.Log(balance)
	assert.NoError(t, err)
	var newbalance = float64(rand.Int31n(10000))
	err = AlterAccountBalance(addr, newbalance)
	assert.NoError(t, err)
	balance2, err := GetDBBalance(addr)
	assert.NoError(t, err)
	t.Log(balance2)
	assert.NotEqualValues(t, balance, balance2)

}

