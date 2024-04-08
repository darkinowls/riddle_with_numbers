package util

import (
	"math/rand"
	"strings"
	"time"
)

const abc = "abcdefghijklmnopqrstuvwxyz"

// first
func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt returns a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(size int) string {
	var sb strings.Builder
	l := len(abc)
	for i := 0; i < size; i++ {
		sb.WriteByte(abc[rand.Intn(l)])
	}
	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomEmail() string {
	return RandomString(6) + "@example.com"
}
