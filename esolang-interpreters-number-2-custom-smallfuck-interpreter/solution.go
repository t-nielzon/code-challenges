package main

func Interpreter(code string, tape string) string {
	t := make([]byte, len(tape))
	for i := 0; i < len(tape); i++ {
		t[i] = tape[i]
	}
	
	ptr := 0
	ip := 0
	
	for ip < len(code) {
		switch code[ip] {
		case '>':
			ptr++
			if ptr >= len(t) {
				return string(t)
			}
		case '<':
			ptr--
			if ptr < 0 {
				return string(t)
			}
		case '*':
			if t[ptr] == '0' {
				t[ptr] = '1'
			} else {
				t[ptr] = '0'
			}
		case '[':
			if t[ptr] == '0' {
				depth := 1
				for ip++; ip < len(code); ip++ {
					if code[ip] == '[' {
						depth++
					} else if code[ip] == ']' {
						depth--
						if depth == 0 {
							break
						}
					}
				}
			}
		case ']':
			if t[ptr] != '0' {
				depth := 1
				for ip--; ip >= 0; ip-- {
					if code[ip] == ']' {
						depth++
					} else if code[ip] == '[' {
						depth--
						if depth == 0 {
							break
						}
					}
				}
				ip--
			}
		}
		ip++
	}
	
	return string(t)
}