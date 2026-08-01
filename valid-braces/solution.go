package main

func ValidBraces(s string) bool {
	stack := []rune{}
	matches := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, ch := range s {
		if closing, isClosing := matches[ch]; isClosing {
			if len(stack) == 0 || stack[len(stack)-1] != closing {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}

	return len(stack) == 0
}