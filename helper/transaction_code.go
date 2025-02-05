package helper

import (
	"fmt"
	"time"

	"math/rand"
)

func GenerateTransactionCode() string {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXY1234567890"
	randomPart := make([]byte, 5)

	for i := range randomPart {
		randomPart[i] = charset[rng.Intn(len(charset))]
	}

	currentDate := time.Now().Format("20060102")

	transactionCode := fmt.Sprintf("CS%s%s", string(randomPart), currentDate)

	return transactionCode
}
