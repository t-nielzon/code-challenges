package main

func DivisibilityRule13(n int64) int64 {
	sequence := []int64{1, 10, 9, 12, 3, 4}
	
	for {
		sum := int64(0)
		temp := n
		seqIdx := 0
		
		if temp == 0 {
			return 0
		}
		
		// Extract digits from right to left and apply the sequence
		for temp > 0 {
			digit := temp % 10
			sum += digit * sequence[seqIdx%len(sequence)]
			temp /= 10
			seqIdx++
		}
		
		// Stop when stationary
		if sum == n {
			return n
		}
		
		n = sum
	}
}