package kata

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func CountOnes(n int) int {
	if n <= 0 {
		return 0
	}
	count := 1
	a := 7
	for i := 2; i <= n; i++ {
		next := a + gcd(i, a)
		if next-a == 1 {
			count++
		}
		a = next
	}
	return count
}

func MaxPn(n int) int {
	if n <= 0 {
		return 0
	}
	seen := make(map[int]bool)
	primes := []int{}
	a := 7
	for i := 2; len(primes) < n; i++ {
		next := a + gcd(i, a)
		d := next - a
		if d != 1 && !seen[d] {
			seen[d] = true
			primes = append(primes, d)
		}
		a = next
	}
	max := primes[0]
	for _, p := range primes {
		if p > max {
			max = p
		}
	}
	return max
}

func AnOverAverage(n int) int {
	if n <= 0 {
		return 0
	}
	sum := 0
	count := 0
	a := 7
	for i := 2; count < n; i++ {
		next := a + gcd(i, a)
		if next-a != 1 {
			sum += next / i
			count++
		}
		a = next
	}
	return sum / n
}