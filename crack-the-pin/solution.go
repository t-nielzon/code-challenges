package main

import (
	"crypto/md5"
	"fmt"
)

func CrackPin(hash string) string {
	for i := 0; i < 100000; i++ {
		pin := fmt.Sprintf("%05d", i)
		h := md5.Sum([]byte(pin))
		hashStr := fmt.Sprintf("%x", h)
		if hashStr == hash {
			return pin
		}
	}
	return ""
}