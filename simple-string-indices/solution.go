package kata

import "errors"

func Solve(s string, idx int) (int, error) {
	if idx < 0 || idx >= len(s) || s[idx] != '(' {
		return -1, errors.New("no opening bracket at the given index")
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
	return -1, errors.New("no matching closing bracket")
}