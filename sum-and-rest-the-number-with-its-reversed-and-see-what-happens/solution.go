package kata

var cache []int

func reverseNum(n int) int {
	rev := 0
	for n > 0 {
		rev = rev*10 + n%10
		n /= 10
	}
	return rev
}

func SumDivRev(n int) int {
	for len(cache) < n {
		start := 10
		if len(cache) > 0 {
			start = cache[len(cache)-1] + 1
		}
		for i := start; ; i++ {
			if i%10 == 0 {
				continue
			}
			rev := reverseNum(i)
			if rev == i {
				continue
			}
			sum := i + rev
			diff := i - rev
			if diff < 0 {
				diff = -diff
			}
			if sum%diff == 0 {
				cache = append(cache, i)
				if len(cache) >= n {
					break
				}
			}
		}
	}
	return cache[n-1]
}