package kata

import "math"

func ListSquared(m, n int) [][2]int {
	result := [][2]int{}
	for num := m; num <= n; num++ {
		sum := sumSquaredDivisors(num)
		sqrt := int(math.Round(math.Sqrt(float64(sum))))
		if sqrt*sqrt == sum {
			result = append(result, [2]int{num, sum})
		}
	}
	return result
}

func sumSquaredDivisors(n int) int {
	sum := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			sum += i * i
			if i != n/i {
				sum += (n / i) * (n / i)
			}
		}
	}
	return sum
}