package kata

import (
	"math"
	"strings"
)

func Code(s string) string {
	l := len(s)
	n := int(math.Ceil(math.Sqrt(float64(l))))
	if n*n < l {
		n++
	}
	padded := s + strings.Repeat("\x0b", n*n-l)
	rows := make([]string, n)
	for i := 0; i < n; i++ {
		rows[i] = padded[i*n : (i+1)*n]
	}
	rotated := make([][]byte, n)
	for i := 0; i < n; i++ {
		rotated[i] = make([]byte, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rotated[j][n-1-i] = rows[i][j]
		}
	}
	out := make([]string, n)
	for i := 0; i < n; i++ {
		out[i] = string(rotated[i])
	}
	return strings.Join(out, "\n")
}

func Decode(s string) string {
	rows := strings.Split(s, "\n")
	n := len(rows)
	rotated := make([][]byte, n)
	for i := 0; i < n; i++ {
		rotated[i] = make([]byte, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rotated[n-1-j][i] = rows[i][j]
		}
	}
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.Write(rotated[i])
	}
	return strings.TrimRight(b.String(), "\x0b")
}