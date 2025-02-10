package helper

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateOTP() string {
	n, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%06d", n.Int64())
}
