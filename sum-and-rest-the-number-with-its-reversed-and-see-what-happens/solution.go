package kata

var sequence []int
var lastChecked = 9

func reverseNum(n int) int {
	rev := 0
	for n > 0 {
		rev = rev*10 + n%10
		n /= 10
	}
	return rev
}

func SumAndRest(n int) int {
	for len(sequence) < n {
		lastChecked++
		i := lastChecked
		// skip numbers ending in 0 — reversed would have a leading zero
		if i%10 == 0 {
			continue
		}
		rev := reverseNum(i)
		// skip palindromes — difference is zero, division undefined
		if rev == i {
			continue
		}
		sum := i + rev
		diff := i - rev
		if diff < 0 {
			diff = -diff
		}
		if sum%diff == 0 {
			sequence = append(sequence, i)
		}
	}
	return sequence[n-1]
}