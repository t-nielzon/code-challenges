package kata

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	if num%2 == 0 {
		return num == 2
	}
	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func PrimeBefAft(n int) [2]int {
	bef := n - 1
	for !isPrime(bef) {
		bef--
	}
	aft := n + 1
	for !isPrime(aft) {
		aft++
	}
	return [2]int{bef, aft}
}