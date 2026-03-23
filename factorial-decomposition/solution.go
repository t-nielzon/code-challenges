package kata

import (
	"fmt"
	"strings"
)

func Decomp(n int) string {
	primes := sieve(n)
	parts := make([]string, 0, len(primes))
	for _, p := range primes {
		exp := 0
		pk := p
		for pk <= n {
			exp += n / pk
			pk *= p
		}
		if exp == 1 {
			parts = append(parts, fmt.Sprintf("%d", p))
		} else {
			parts = append(parts, fmt.Sprintf("%d^%d", p, exp))
		}
	}
	return strings.Join(parts, " * ")
}

func sieve(n int) []int {
	if n < 2 {
		return nil
	}
	mark := make([]bool, n+1)
	for i := 2; i*i <= n; i++ {
		if !mark[i] {
			for j := i * i; j <= n; j += i {
				mark[j] = true
			}
		}
	}
	primes := make([]int, 0)
	for i := 2; i <= n; i++ {
		if !mark[i] {
			primes = append(primes, i)
		}
	}
	return primes
}