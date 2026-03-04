package kata

func SumMultTriangNumb(n, m int) int {
	triangulars := make([]int, n)
	for i := 1; i <= n; i++ {
		triangulars[i-1] = i * (i + 1) / 2
	}

	lcmVal := triangulars[0]
	for i := 1; i < n; i++ {
		lcmVal = lcm(lcmVal, triangulars[i])
	}

	return lcmVal * m * (m + 1) / 2
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}