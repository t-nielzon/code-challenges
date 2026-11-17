package main

/*
1, 246, 2, 123, 3, 82, 6, 41 are the divisors of number 246.

Squaring these divisors we get: 1, 60516, 4, 15129, 9, 6724, 36, 1681.

The sum of these squares is 84100 which is 290 * 290.

Task:
Find all integers between m and n (m and n are integers with 1 <= m <= n)
such that the sum of their squared divisors is itself a square.

Return an array of subarrays with two elements:
- first: the number the squared divisors of which is a square
- second: the sum of the squared divisors

Example:
m =  1, n = 250 --> [[1, 1], [42, 2500], [246, 84100]]
m = 42, n = 250 --> [[42, 2500], [246, 84100]]
*/

func SumOfSquaredDivisors(m, n int) [][]int {
	var result [][]int

	for num := m; num <= n; num++ {
		sumOfSquares := 0

		for i := 1; i*i <= num; i++ {
			if num%i == 0 {
				sumOfSquares += i * i
				if i != num/i {
					sumOfSquares += (num/i) * (num/i)
				}
			}
		}

		if isPerfectSquare(sumOfSquares) {
			result = append(result, []int{num, sumOfSquares})
		}
	}

	return result
}

func isPerfectSquare(n int) bool {
	if n < 0 {
		return false
	}
	if n == 0 {
		return true
	}

	x := n
	y := (x + 1) / 2
	for y < x {
		x = y
		y = (x + n/x) / 2
	}
	return x*x == n
}