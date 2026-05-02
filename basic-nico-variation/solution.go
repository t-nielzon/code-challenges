package kata

import (
	"sort"
	"strings"
)

func Nico(key, message string) string {
	n := len(key)
	if n == 0 {
		return message
	}
	sorted := []byte(key)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })

	labels := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if sorted[j] == key[i] {
				labels[i] = j + 1
				break
			}
		}
	}

	pad := (n - len(message)%n) % n
	msg := message + strings.Repeat(" ", pad)
	rows := len(msg) / n

	order := make([]int, n)
	for k := 1; k <= n; k++ {
		for idx := 0; idx < n; idx++ {
			if labels[idx] == k {
				order[k-1] = idx
				break
			}
		}
	}

	out := make([]byte, 0, len(msg))
	for r := 0; r < rows; r++ {
		for k := 0; k < n; k++ {
			out = append(out, msg[r*n+order[k]])
		}
	}
	return string(out)
}