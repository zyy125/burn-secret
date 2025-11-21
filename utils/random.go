package utils

import (
	"math/rand"
	"time"
)

func GetRandomID(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seedRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)

	for i:= range b{
		b[i] = charset[seedRand.Intn(len(charset))]
	}

	return string(b)
}