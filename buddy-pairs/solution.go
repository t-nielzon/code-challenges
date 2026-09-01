// solution.go
package main

func buddy(start, limit int) []int {
	for n := start; n <= limit; n++ {
		sn := sumProperDivisors(n)
		m := sn - 1
		
		if m > n && sumProperDivisors(m) == n+1 {
			return []int{n, m}
		}
	}
	return nil
}

func sumProperDivisors(n int) int {
	if n <= 1 {
		return 0
	}
	
	sum := 1 // 1 is always a proper divisor for n > 1
	
	// find divisors up to sqrt(n)
	i := 2
	for i*i <= n {
		if n%i == 0 {
			sum += i
			if i*i != n {
				sum += n / i
			}
		}
		i++
	}
	
	return sum
}