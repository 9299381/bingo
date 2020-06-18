package util

import (
	"math/rand"
	"time"
)

func RandString(length int, opt ...string) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if opt != nil {
		if opt[0] == "0" {
			str = "0123456789"
		} else if opt[0] == "a" {
			str = "abcdefghijklmnopqrstuvwxyz"
		} else if opt[0] == "A" {
			str = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		} else if opt[0] == "aA" {
			str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		}
	}
	bytes := []byte(str)
	bytesLen := len(bytes)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(bytesLen)])
	}
	return string(result)
}
