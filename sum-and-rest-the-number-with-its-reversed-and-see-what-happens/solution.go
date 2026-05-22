package kata

func reverseInt(n int) int {
	r := 0
	for n > 0 {
		r = r*10 + n%10
		n /= 10
	}
	return r
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var cache []int

func computeUpTo(n int) {
	start := 10
	if len(cache) > 0 {
		start = cache[len(cache)-1] + 1
	}
	for k := start; len(cache) < n; k++ {
		if k%10 == 0 {
			continue
		}
		r := reverseInt(k)
		if r == k {
			continue
		}
		diff := abs(k - r)
		if (k+r)%diff == 0 {
			cache = append(cache, k)
		}
	}
}

func SumRestReversed(n int) int {
	if n <= 0 {
		return 0
	}
	if len(cache) < n {
		computeUpTo(n)
	}
	return cache[n-1]
}