package util

import (
	"fmt"
	"math/big"
	"strings"
)


var Unit = map[int]float64{
	1:  1e1,
	2:  1e2,
	3:  1e3,
	4:  1e4,
	5:  1e5,
	6:  1e6,
	7:  1e7,
	8:  1e8,
	9:  1e9,
	10: 1e10,
	11: 1e11,
	12: 1e12,
	13: 1e13,
	14: 1e14,
	15: 1e15,
	16: 1e16,
	17: 1e17,
	18: 1e18,
}

//去掉末尾的0和小数点
func TrimZeroAndDot(s string) string {
	if strings.Contains(s, ".") {
		trimZeroStr := strings.TrimRight(s, "0")
		trimDotStr := strings.TrimRight(trimZeroStr, ".")
		return trimDotStr
	}
	return s
}

func CaculBalance(bn int64) string {
	bf := big.NewFloat(float64(bn))
	bf = bf.Quo(bf, big.NewFloat(1e8))
	f64, _ := bf.Float64()
	return fmt.Sprintf("%.8f", f64)
}




// 按照decimal精度转换为big.Int类型数据
func ToBig(amount float64, decimal int) *big.Int{
	amount += 1/ (Unit[decimal]*10)
	bf:= big.NewFloat(amount)
	bf= bf.Mul(bf,big.NewFloat(Unit[decimal]))
	bn,_ := bf.Int(nil)
	return bn
}