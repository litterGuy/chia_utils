package chia_utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
)

func GetCoinIdInfo(amountJ *big.Int, parentCoinInfo string, puzzleHash string) (string, error) {
	amountHex := fixPreLength(amountJ)
	if strings.HasPrefix(parentCoinInfo, "0x") {
		parentCoinInfo = strings.TrimLeft(parentCoinInfo, "0x")
	}
	if strings.HasPrefix(puzzleHash, "0x") {
		puzzleHash = strings.TrimLeft(puzzleHash, "0x")
	}
	tmpid, err := hex.DecodeString(parentCoinInfo)
	if err != nil {
		return "", err
	}
	tmppuzzle, err := hex.DecodeString(puzzleHash)
	if err != nil {
		return "", err
	}
	tmpAmount, err := hex.DecodeString(amountHex)
	if err != nil {
		return "", err
	}

	tmp := append(tmpid, tmppuzzle...)
	tmp = append(tmp, tmpAmount...)

	rt := sha256.Sum256(tmp)
	return "0x" + hex.EncodeToString(rt[:]), nil
}

func fixPreLength(amount *big.Int) string {
	amountHex := fmt.Sprintf("%x", amount)
	byteCount := (len(fmt.Sprintf("%b", amount)) + 8) >> 3
	fixLen := byteCount * 2
	tmp := strings.Repeat("0", fixLen-1) + amountHex
	amountHex = tmp[len(tmp)-fixLen:]
	if len(amountHex)%2 == 1 {
		amountHex = "0" + amountHex
	}
	return amountHex
}
