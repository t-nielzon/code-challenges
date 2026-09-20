package main

func PowerOfDigitSum(n int) int {
	count := 0
	num := 10

	for {
		if isPowerOfDigitSum(num) {
			count++
			if count == n {
				return num
			}
		}
		num++
	}
}

func isPowerOfDigitSum(num int) bool {
	digitSum := 0
	temp := num
	for temp > 0 {
		digitSum += temp % 10
		temp /= 10
	}

	if digitSum == 0 {
		return false
	}

	power := digitSum
	for power < num {
		power *= digitSum
	}

	return power == num
}