package kata

import "math"

func ListSquared(m, n int) [][]int {
	var result [][]int
	for i := m; i <= n; i++ {
		sum := sumOfSquaredDivisors(i)
		root := int(math.Sqrt(float64(sum)))
		if root*root == sum {
			result = append(result, []int{i, sum})
		}
	}
	return result
}

func sumOfSquaredDivisors(n int) int {
	sum := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			sum += i * i
			j := n / i
			if j != i {
				sum += j * j
			}
		}
	}
	return sum
}