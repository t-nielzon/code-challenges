package kata

func PrimeBefAft(n int) [2]int {
	return [2]int{primeBefore(n), primeAfter(n)}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primeBefore(n int) int {
	for i := n - 1; i >= 2; i-- {
		if isPrime(i) {
			return i
		}
	}
	return -1
}

func primeAfter(n int) int {
	for i := n + 1; ; i++ {
		if isPrime(i) {
			return i
		}
	}
}