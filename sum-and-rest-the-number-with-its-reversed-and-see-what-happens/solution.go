package kata

var sumDivCache = []int{}

func reverse(x int) int {
	r := 0
	for x > 0 {
		r = r*10 + x%10
		x /= 10
	}
	return r
}

func SumDivSequence(n int) int {
	for len(sumDivCache) < n {
		// candidate numbering continues from the last found term
		next := 1
		if len(sumDivCache) > 0 {
			next = sumDivCache[len(sumDivCache)-1] + 1
		}
		for num := next; ; num++ {
			// numbers ending in 0 reverse with a leading zero; skip them
			if num%10 == 0 {
				continue
			}
			rev := reverse(num)
			diff := num - rev
			if diff < 0 {
				diff = -diff
			}
			if diff == 0 {
				continue
			}
			if (num+rev)%diff == 0 {
				sumDivCache = append(sumDivCache, num)
				break
			}
		}
	}
	return sumDivCache[n-1]
}