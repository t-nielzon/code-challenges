package main

func Xbonacci(signature []int, n int) []int {
	result := make([]int, 0, n)
	
	// Add the signature elements (or up to n if signature is shorter)
	for i := 0; i < len(signature) && i < n; i++ {
		result = append(result, signature[i])
	}
	
	// Generate remaining elements
	for len(result) < n {
		sum := 0
		// Sum the last len(signature) elements
		for i := len(result) - len(signature); i < len(result); i++ {
			sum += result[i]
		}
		result = append(result, sum)
	}
	
	return result
}