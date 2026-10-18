package main

func DivisibilityBy13(n int) int {
	sequence := []int{1, 10, 9, 12, 3, 4}
	
	for {
		sum := 0
		temp := n
		index := 0
		
		// extract digits from right to left and multiply by sequence
		for temp > 0 {
			digit := temp % 10
			sum += digit * sequence[index%len(sequence)]
			temp /= 10
			index++
		}
		
		// check if stationary
		if sum == n {
			return n
		}
		n = sum
	}
}