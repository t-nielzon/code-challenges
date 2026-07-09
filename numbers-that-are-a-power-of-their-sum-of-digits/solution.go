package main

func PowerSumDigits(n int) int {
	count := 0
	num := 1

	for {
		// calculate sum of digits
		sumDigits := 0
		temp := num
		for temp > 0 {
			sumDigits += temp % 10
			temp /= 10
		}

		// check if any power of sumDigits equals num
		for power := 1; ; power++ {
			result := 1
			for i := 0; i < power; i++ {
				result *= sumDigits
				if result > num {
					break
				}
			}

			if result == num {
				count++
				if count == n {
					return num
				}
				break
			}

			if result > num {
				break
			}
		}

		num++
	}
}