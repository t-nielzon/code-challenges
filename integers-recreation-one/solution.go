package kata

import "math"

func ListSquared(m, n int) [][]int {
	result := [][]int{}

	for num := m; num <= n; num++ {
		sumSquares := sumOfSquaredDivisors(num)
		if isPerfectSquare(sumSquares) {
			result = append(result, []int{num, sumSquares})
		}
	}

	return result
}

func sumOfSquaredDivisors(n int) int {
	sum := 0
	sqrt := int(math.Sqrt(float64(n)))

	for i := 1; i <= sqrt; i++ {
		if n%i == 0 {
			sum += i * i
			other := n / i
			if other != i {
				sum += other * other
			}
		}
	}

	return sum
}

func isPerfectSquare(n int) bool {
	if n < 0 {
		return false
	}
	sqrt := int(math.Sqrt(float64(n)))
	return sqrt*sqrt == n
}