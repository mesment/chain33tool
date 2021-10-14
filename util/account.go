package util

import (
	"errors"
	"fmt"
	dbm "github.com/33cn/chain33/common/db"
	"github.com/33cn/chain33/rpc/jsonclient"
	rpctypes "github.com/33cn/chain33/rpc/types"
	_ "github.com/33cn/chain33/system"
	"github.com/33cn/chain33/types"
	"github.com/33cn/plugin/plugin/store/kvmvccmavl"
	_ "github.com/bityuan/bityuan/plugin"
)



var (
	db dbm.DB
	coinsDBPrefix = ".-mvcc-.d.mavl-coins-bty"

	lastPrefix    = ".-mvcc-.l.mavl-coins-bty"
	ErrNotFount = errors.New("not found")
)

func InitializeDB(chain33Cfg *types.Chain33Config) dbm.DB {
	sCfg := chain33Cfg.GetModuleConfig().Store
	fmt.Println(sCfg)
	store := kvmvccmavl.New(sCfg, nil, chain33Cfg).(*kvmvccmavl.KVmMavlStore)
	db = store.GetDB()
	return db
}

func CloseDB()  {
	db.Close()
}


func GetDBBalance(addr string)  (string, error){
	b, err := getLocalBalance(addr)
	if err != nil {
		return "0", err
	}
	balance := CaculBalance(b)
	return balance, nil
}

func getLocalBalance(addr string)  (int64, error){
	addrPrefix := fmt.Sprintf("%s-%s", coinsDBPrefix, addr)
	//fmt.Printf("db prefix:%s\n", addrPrefix)
	kv, err := GetKV(addrPrefix)
	if err == ErrNotFount {
		fmt.Printf("%s not found\n", addrPrefix)
		return 0, nil
	}

	var acc types.Account
	err = types.Decode(kv.GetValue(), &acc)
	if err != nil {
		fmt.Printf("Decode value err:%v\n",err)
		return 0, err
	}
	//fmt.Printf("addr:%s, key:%s,  local balance:%s\n",  addr, kv.GetKey(), CaculBalance(acc.Balance))
	return acc.Balance, nil
}

func AlterAccountBalance(addr string, newBalance float64) error {
	var (
		key string
	    acc types.Account
		updateLastKey bool
	)

	addrPrefix := fmt.Sprintf("%s-%s", coinsDBPrefix, addr)
	//fmt.Printf("db prefix:%s\n", addrPrefix)
	kv, err := GetKV(addrPrefix)
	if err != nil && err != ErrNotFount{
		fmt.Println(err)
		return err
	}
	if err == ErrNotFount {
		key = fmt.Sprintf("%s-%s.%020d",coinsDBPrefix, addr, 0)
	} else {
		updateLastKey = true

		key = string(kv.GetKey())
		err = types.Decode(kv.GetValue(), &acc)
		if err != nil {
			fmt.Printf("Decode value err:%v\n",err)
			return err
		}

	}
	//fmt.Printf("addr:%s, key:%s,local balance:%s\n",addr, key,  CaculBalance(acc.Balance))

	newValue := ToBig(newBalance, 8).Int64()
	acc.Balance = newValue

	if acc.Addr == "" {
		acc.Addr = addr
	}
	accBytes := types.Encode(&acc)
	err = db.Set([]byte(key), accBytes)
	if err != nil {
		fmt.Println("AlterAccountBalance set new balance err:", err)
		return err
	}
	if updateLastKey {
		lastKey := fmt.Sprintf("%s-%s", lastPrefix, addr)
		//fmt.Println("lastKey:",lastKey)
		err = db.Set([]byte(lastKey), accBytes)
		if err != nil {
			fmt.Println("AlterAccountBalance set new balance err:", err)
			return err
		}
	}
	return nil
}

func GetKV(prefix string ) (*types.KeyValue, error)  {
	it := dbm.NewListHelper(db)
	list := it.IteratorScanFromLast([]byte(prefix), 1, dbm.ListWithKey)
	var v types.KeyValue
	if len(list) > 0 {
		err := types.Decode(list[0],&v)
		if err != nil {
			fmt.Printf("GetKV Decode list err:%v\n", err)
			return nil, err
		} else {
			//fmt.Printf("keyvalue: key:%s, value:%s\n", v.GetKey(), v.GetValue())
			return &v, nil
		}
	}
	return nil, ErrNotFount
}



func GetBalance(addr string,rpc_addr string) (string, error) {
	jsonCli, err := jsonclient.NewJSONClient(rpc_addr)
	if err != nil {
		fmt.Println("GetBalance NewJSONClient err:", err)
		return "0", err
	}
	params := types.ReqBalance{
		Addresses: []string{addr},
		Execer:    "coins",
	}
	var result []*rpctypes.Account
	err = jsonCli.Call("Chain33.GetBalance", params, &result)
	if err != nil {
		fmt.Println("GetBalance err:", err)
		return "0",err
	}
	balance := TrimZeroAndDot(CaculBalance(result[0].Balance))
	return balance, nil
}