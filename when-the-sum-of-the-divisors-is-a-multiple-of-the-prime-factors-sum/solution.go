package kata

func primeFactorSum(n int) int {
	sum := 0
	for d := 2; d*d <= n; d++ {
		for n%d == 0 {
			sum += d
			n /= d
		}
	}
	if n > 1 {
		sum += n
	}
	return sum
}

func divisorSum(n int) int {
	sum := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			if i != n/i {
				sum += n / i
			}
		}
	}
	return sum
}

func DsMultofPfs(nMin, nMax int) []int {
	result := []int{}
	for n := nMin; n <= nMax; n++ {
		if n < 2 {
			continue
		}
		pfs := primeFactorSum(n)
		if pfs == 0 {
			continue
		}
		ds := divisorSum(n)
		if ds%pfs == 0 {
			result = append(result, n)
		}
	}
	return result
}