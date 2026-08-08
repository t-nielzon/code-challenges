package main

func DivisibleBy13(n int) int {
	remainders := []int{1, 10, 9, 12, 3, 4}
	
	for {
		sum := 0
		digits := getDigits(n)
		
		for i, digit := range digits {
			sum += digit * remainders[i%len(remainders)]
		}
		
		if sum == n {
			return n
		}
		n = sum
	}
}

func getDigits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	return digits
}