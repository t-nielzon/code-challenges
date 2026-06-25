package kata

import (
	"strconv"
	"strings"
)

func decomp(n int) string {
	// sieve of eratosthenes for primes up to n
	sieve := make([]bool, n+1)
	var parts []string

	for p := 2; p <= n; p++ {
		if sieve[p] {
			continue
		}
		for m := p * p; m <= n; m += p {
			sieve[m] = true
		}

		// legendre's formula: exponent of prime p in n! is sum of floor(n/p^k)
		exp := 0
		for pk := p; pk <= n; pk *= p {
			exp += n / pk
		}

		if exp == 1 {
			parts = append(parts, strconv.Itoa(p))
		} else {
			parts = append(parts, strconv.Itoa(p)+"^"+strconv.Itoa(exp))
		}
	}

	return strings.Join(parts, " * ")
}