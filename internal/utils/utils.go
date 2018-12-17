package utils

import (
	"encoding/base64"
	"fmt"
	"os"
)

func UintToBase64(val uint) string {
	data := []byte(fmt.Sprint(val))
	return base64.StdEncoding.EncodeToString(data)
}

func Getenv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
