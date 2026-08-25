package main

import (
	"fmt"
	"strings"
)

func decomp(n int) string {
	if n < 2 {
		return "1"
	}
	
	primes := sieve(n)
	var result []string
	
	for _, p := range primes {
		count := 0
		power := p
		for power <= n {
			count += n / power
			if power > n/p {
				break
			}
			power *= p
		}
		
		if count == 1 {
			result = append(result, fmt.Sprintf("%d", p))
		} else {
			result = append(result, fmt.Sprintf("%d^%d", p, count))
		}
	}
	
	return strings.Join(result, " * ")
}

func sieve(n int) []int {
	if n < 2 {
		return []int{}
	}
	
	is_prime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		is_prime[i] = true
	}
	
	for i := 2; i*i <= n; i++ {
		if is_prime[i] {
			for j := i * i; j <= n; j += i {
				is_prime[j] = false
			}
		}
	}
	
	var primes []int
	for i := 2; i <= n; i++ {
		if is_prime[i] {
			primes = append(primes, i)
		}
	}
	
	return primes
}