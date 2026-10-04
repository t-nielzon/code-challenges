package main

import (
	"strconv"
)

func DigitsPower(n, p int) int {
	digits := strconv.Itoa(n)
	var sum int64 = 0

	for i, digit := range digits {
		d := int64(digit - '0')
		power := int64(1)
		for j := 0; j < p+i; j++ {
			power *= d
		}
		sum += power
	}

	if sum%int64(n) == 0 {
		return int(sum / int64(n))
	}
	return -1
}