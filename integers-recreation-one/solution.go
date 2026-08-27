package main

import "math"

func SumOfSquaredDivisors(m, n int) [][]int {
	var result [][]int
	
	for num := m; num <= n; num++ {
		sumOfSquares := 0
		
		// find all divisors and sum their squares
		for i := 1; i*i <= num; i++ {
			if num%i == 0 {
				sumOfSquares += i * i
				if i != num/i {
					sumOfSquares += (num/i) * (num/i)
				}
			}
		}
		
		// check if sum is a perfect square
		sqrt := int(math.Sqrt(float64(sumOfSquares)))
		if sqrt*sqrt == sumOfSquares {
			result = append(result, []int{num, sumOfSquares})
		}
	}
	
	return result
}