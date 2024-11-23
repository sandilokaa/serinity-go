package helper

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func APIResponseDelete(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

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

func GenerateStateOauth() string {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	stateLength := 32
	state := make([]byte, stateLength)

	for i := range state {
		state[i] = charset[rng.Intn(len(charset))]
	}

	return string(state)
}

func GenerateRandomCodeOauth(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	code := make([]byte, length)
	for i := range code {
		code[i] = charset[rng.Intn(len(charset))]
	}

	return string(code)
}
