package helper

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

// 1. create object response
type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"` // why interface{}? bcoz value of the data can change
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

// 2. Mapping value of response from handler
func ApiResponse(message string, code int, status string, data interface{}) Response {
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

func FormatValidationError(err error) []string {
	// setup validation
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func TimeNowMilli() string {
	timeMilli := time.Now().UnixNano() / int64(time.Millisecond)
	return strconv.Itoa(int(timeMilli))
}

func GenerateCodeTransaction(userID int) string {
	code := fmt.Sprintf("TRX-%s-%s", strconv.Itoa(userID), TimeNowMilli())
	return code
}

func GetENV(key, fallback string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
