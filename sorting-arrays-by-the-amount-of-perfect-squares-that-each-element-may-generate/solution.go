package kata

import (
	"math"
	"sort"
	"strconv"
)

func SortByPerfSq(arr []int) []int {
	counts := make(map[int]int, len(arr))
	for _, n := range arr {
		counts[n] = countPerfectSquares(n)
	}
	result := make([]int, len(arr))
	copy(result, arr)
	sort.SliceStable(result, func(i, j int) bool {
		ci, cj := counts[result[i]], counts[result[j]]
		if ci != cj {
			return ci > cj
		}
		return result[i] < result[j]
	})
	return result
}

func countPerfectSquares(n int) int {
	digits := strconv.Itoa(n)
	seen := make(map[string]bool)
	count := 0
	permute([]byte(digits), 0, seen, &count, len(digits))
	return count
}

func permute(digits []byte, idx int, seen map[string]bool, count *int, length int) {
	if idx == length {
		s := string(digits)
		if seen[s] {
			return
		}
		seen[s] = true
		if digits[0] == '0' {
			return
		}
		n, _ := strconv.Atoi(s)
		sq := int(math.Round(math.Sqrt(float64(n))))
		if sq*sq == n {
			*count++
		}
		return
	}
	for i := idx; i < length; i++ {
		digits[idx], digits[i] = digits[i], digits[idx]
		permute(digits, idx+1, seen, count, length)
		digits[idx], digits[i] = digits[i], digits[idx]
	}
}