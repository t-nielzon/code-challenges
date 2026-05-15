package kata

import "fmt"

func Solve(s string, idx int) (int, error) {
	if idx < 0 || idx >= len(s) || s[idx] != '(' {
		return -1, fmt.Errorf("no opening bracket at index %d", idx)
	}
	depth := 0
	for i := idx; i < len(s); i++ {
		switch s[i] {
		case '(':
			depth++
		case ')':
			depth--
			if depth == 0 {
				return i, nil
			}
		}
	}
	return -1, fmt.Errorf("no matching closing bracket")
}