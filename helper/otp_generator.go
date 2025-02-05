package helper

import (
	"crypto/rand"
	"fmt"
	"log"
)

func GenerateOTP() string {
	otp := make([]byte, 3)
	_, err := rand.Read(otp)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%06d", int(otp[0])%1000000)
}
