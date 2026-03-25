package kata

import (
	"fmt"
	"strings"
)

func Collatz(n int) string {
	parts := []string{fmt.Sprintf("%d", n)}
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		parts = append(parts, fmt.Sprintf("%d", n))
	}
	return strings.Join(parts, "->")
}