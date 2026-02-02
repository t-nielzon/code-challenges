package kata

import (
	"fmt"
	"strings"
)

func CoffeeLimits(birthDate int) [2]int {
	const (
		CAFE  = 0xCAFE  // 51966
		DECAF = 0xDECAF // 912559
		limit = 5000
	)

	containsDead := func(n int) bool {
		hex := fmt.Sprintf("%X", n)
		return strings.Contains(hex, "DEAD")
	}

	findLimit := func(coffeeValue int) int {
		h := birthDate
		for cups := 1; cups <= limit; cups++ {
			h += coffeeValue
			if containsDead(h) {
				return cups
			}
		}
		return 0
	}

	regularLimit := findLimit(CAFE)
	decafLimit := findLimit(DECAF)

	return [2]int{regularLimit, decafLimit}
}