package miniredis

import (
	"math/big"
)

func parseFloat(s string) (*big.Float, error) {
	f, _, err := big.ParseFloat(s, 0, 256, 0)
	return f, err
}

func newZero() *big.Float {
	v := big.NewFloat(0.0)
	v.SetPrec(256)
	return v
}
