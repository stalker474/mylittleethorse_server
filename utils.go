package main

import (
	"math"
	"math/big"
)

// ToBytes32 convert a string to bytes32 (by filling 0)
func ToBytes32(value string) [32]byte {
	var array [32]byte
	copy(array[:], value)
	return array
}

// FromBytes32 convert a bytes 32 to null terminated string
func FromBytes32(value [32]byte) string {
	return string(value[:])
}

// WeiToEth convert wei to eth
func WeiToEth(value *big.Int) (res float32) {
	f := new(big.Float)
	f2 := big.NewFloat(math.Pow(10, (18)))
	f.SetInt(value)
	z := new(big.Float).Quo(f, f2)

	res, _ = z.Float32()
	return res
}
