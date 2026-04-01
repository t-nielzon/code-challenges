package kata

func PowerSumDigTerm(n int) int {
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

func digitSum(n int) int {
	s := 0
	for n > 0 {
		s += n % 10
		n /= 10
	}
	return s
}

func isPowerOfDigitSum(n int) bool {
	s := digitSum(n)
	if s < 2 {
		return false
	}
	power := s * s
	for power <= n {
		if power == n {
			return true
		}
		power *= s
	}
	return false
}