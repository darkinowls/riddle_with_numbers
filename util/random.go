package util

import (
	"math/rand"
	"strings"
	"time"
)

const abc = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// first
func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(size int) string {
	var sb strings.Builder
	l := len(abc)
	for i := 0; i < size; i++ {
		sb.WriteByte(abc[rand.Intn(l)])
	}
	return sb.String()
}

func RandomEmail() string {
	return RandomString(30) + "@example.com"
}
