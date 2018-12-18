package utils

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/gosimple/slug"
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

func ToTagSlice(items []string) []string {
	m := make(map[string]bool)
	o := make([]string, 0, len(items))
	for _, item := range items {
		tag := slug.Make(item)
		if _, ok := m[tag]; !ok {
			m[tag] = true
			o = append(o, tag)
		}
	}
	return o
}
