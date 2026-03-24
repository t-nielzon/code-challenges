package kata

func Interpreter(code, tape string) string {
	t := []byte(tape)
	ptr := 0
	ip := 0

	for ip < len(code) && ptr >= 0 && ptr < len(t) {
		switch code[ip] {
		case '>':
			ptr++
		case '<':
			ptr--
		case '*':
			if t[ptr] == '0' {
				t[ptr] = '1'
			} else {
				t[ptr] = '0'
			}
		case '[':
			if t[ptr] == '0' {
				depth := 1
				for depth > 0 {
					ip++
					if code[ip] == '[' {
						depth++
					} else if code[ip] == ']' {
						depth--
					}
				}
			}
		case ']':
			if t[ptr] != '0' {
				depth := 1
				for depth > 0 {
					ip--
					if code[ip] == ']' {
						depth++
					} else if code[ip] == '[' {
						depth--
					}
				}
			}
		}
		ip++
	}

	return string(t)
}