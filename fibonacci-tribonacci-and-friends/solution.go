package main

func Xbonacci(signature []int, n int) []int {
	if n == 0 {
		return []int{}
	}

	result := make([]int, n)

	// Copy signature into result (up to n elements)
	for i := 0; i < len(signature) && i < n; i++ {
		result[i] = signature[i]
	}

	// Generate remaining elements
	sigLen := len(signature)
	for i := sigLen; i < n; i++ {
		sum := 0
		for j := i - sigLen; j < i; j++ {
			sum += result[j]
		}
		result[i] = sum
	}

	return result
}