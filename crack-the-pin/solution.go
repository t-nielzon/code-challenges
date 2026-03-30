package kata

import (
	"crypto/md5"
	"fmt"
)

func Crack(hash string) string {
	for i := 0; i < 100000; i++ {
		pin := fmt.Sprintf("%05d", i)
		h := fmt.Sprintf("%x", md5.Sum([]byte(pin)))
		if h == hash {
			return pin
		}
	}
	return ""
}