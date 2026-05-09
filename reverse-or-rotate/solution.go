package kata

func RevRot(s string, n int) string {
	if n <= 0 || s == "" || n > len(s) {
		return ""
	}
	result := make([]byte, 0, len(s))
	for i := 0; i+n <= len(s); i += n {
		chunk := s[i : i+n]
		sum := 0
		for j := 0; j < n; j++ {
			sum += int(chunk[j] - '0')
		}
		if sum%2 == 0 {
			for j := n - 1; j >= 0; j-- {
				result = append(result, chunk[j])
			}
		} else {
			result = append(result, chunk[1:]...)
			result = append(result, chunk[0])
		}
	}
	return string(result)
}