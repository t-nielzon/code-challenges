package main

func Interpreter(code string, tape string) string {
	bits := make([]int, len(tape))
	for i, c := range tape {
		bits[i] = int(c - '0')
	}

	pointer := 0
	ip := 0

	for ip < len(code) {
		cmd := code[ip]

		switch cmd {
		case '>':
			pointer++
			if pointer >= len(bits) {
				break
			}
			ip++
		case '<':
			pointer--
			if pointer < 0 {
				break
			}
			ip++
		case '*':
			bits[pointer] = 1 - bits[pointer]
			ip++
		case '[':
			if bits[pointer] == 0 {
				depth := 1
				for ip++; ip < len(code) && depth > 0; ip++ {
					if code[ip] == '[' {
						depth++
					} else if code[ip] == ']' {
						depth--
					}
				}
			} else {
				ip++
			}
		case ']':
			if bits[pointer] != 0 {
				depth := 1
				for ip--; ip >= 0 && depth > 0; ip-- {
					if code[ip] == ']' {
						depth++
					} else if code[ip] == '[' {
						depth--
					}
				}
				ip++
			} else {
				ip++
			}
		default:
			ip++
		}
	}

	result := ""
	for _, b := range bits {
		result += string(byte(b + int('0')))
	}
	return result
}