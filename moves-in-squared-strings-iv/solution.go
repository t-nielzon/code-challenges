package kata

import "strings"

func Diag2Sym(s string) string {
	lines := strings.Split(s, "\n")
	n := len(lines)
	result := make([]string, n)
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = lines[n-1-j][n-1-i]
		}
		result[i] = string(row)
	}
	return strings.Join(result, "\n")
}

func Rot90Counter(s string) string {
	lines := strings.Split(s, "\n")
	n := len(lines)
	result := make([]string, n)
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = lines[j][n-1-i]
		}
		result[i] = string(row)
	}
	return strings.Join(result, "\n")
}

func SelfieDiag2Counterclock(s string) string {
	orig := strings.Split(s, "\n")
	d := strings.Split(Diag2Sym(s), "\n")
	r := strings.Split(Rot90Counter(s), "\n")
	result := make([]string, len(orig))
	for i := range orig {
		result[i] = orig[i] + "|" + d[i] + "|" + r[i]
	}
	return strings.Join(result, "\n")
}

func Oper(fct func(string) string, s string) string {
	return fct(s)
}