package kata

import (
	"fmt"
	"strings"
)

func Decomp(n int) string {
	// Generate all primes up to n using Sieve of Eratosthenes
	primes := sieve(n)
	
	var parts []string
	
	for _, p := range primes {
		// Count how many times prime p divides n!
		// Using Legendre's formula: sum of floor(n/p^k) for k=1,2,3,...
		exp := 0
		pk := p
		for pk <= n {
			exp += n / pk
			pk *= p
		}
		
		if exp > 0 {
			if exp == 1 {
				parts = append(parts, fmt.Sprintf("%d", p))
			} else {
				parts = append(parts, fmt.Sprintf("%d^%d", p, exp))
			}
		}
	}
	
	return strings.Join(parts, " * ")
}

// sieve returns all primes up to and including n
func sieve(n int) []int {
	if n < 2 {
		return []int{}
	}
	
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	
	var primes []int
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	
	return primes
}