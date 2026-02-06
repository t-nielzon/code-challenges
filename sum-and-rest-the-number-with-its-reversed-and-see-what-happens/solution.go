package kata

var cache []int64

func init() {
	cache = buildSequence(65)
}

func reverseNum(n int64) int64 {
	var rev int64
	for n > 0 {
		rev = rev*10 + n%10
		n /= 10
	}
	return rev
}

func buildSequence(limit int) []int64 {
	seq := make([]int64, 0, limit)
	var num int64
	for num = 10; len(seq) < limit; num++ {
		// skip numbers ending in 0 (reversed would have leading zero)
		if num%10 == 0 {
			continue
		}
		rev := reverseNum(num)
		// skip palindromes (difference would be 0)
		if rev == num {
			continue
		}
		sum := num + rev
		diff := num - rev
		if diff < 0 {
			diff = -diff
		}
		if sum%diff == 0 {
			seq = append(seq, num)
		}
	}
	return seq
}

func SumAndRest(n int) int64 {
	if n < 1 || n > len(cache) {
		return 0
	}
	return cache[n-1]
}