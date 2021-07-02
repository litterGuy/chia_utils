package chia_utils

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestGetCoinIdInfo(t *testing.T) {
	amount, _ := decimal.NewFromString("0.81")
	precision := decimal.NewFromInt(1000000000000)
	amountReq := amount.Mul(precision).BigInt()
	parentCoinInfo := "0x6a3860aaf9f4a4d22f0907cfbdb3ee29e138a963381fa2cad909a0c967f7e110"
	puzzleHash := "0xf60182177142f0ace5b3ecd03b11a571817f5411656a4a542e0011f1a0d5e584"

	rt, err := GetCoinIdInfo(amountReq, parentCoinInfo, puzzleHash)
	if err != nil {
		t.Fatal(err)
	}
	println(rt)
}
