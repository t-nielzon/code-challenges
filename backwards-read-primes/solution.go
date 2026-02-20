package kata

func BackwardsPrime(start, stop int) []int {
	var result []int
	for n := start; n <= stop; n++ {
		if isPrime(n) {
			rev := reverse(n)
			if rev != n && isPrime(rev) {
				result = append(result, n)
			}
		}
	}
	return result
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n < 4 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func reverse(n int) int {
	rev := 0
	for n > 0 {
		rev = rev*10 + n%10
		n /= 10
	}
	return rev
}