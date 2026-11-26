package kata

import "strings"

func Diag2Sym(s string) string {
	grid := stringToGrid(s)
	n := len(grid)
	result := make([][]rune, n)
	for i := 0; i < n; i++ {
		result[i] = make([]rune, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[i][j] = grid[n-1-j][n-1-i]
		}
	}

	return gridToString(result)
}

func Rot90Counter(s string) string {
	grid := stringToGrid(s)
	n := len(grid)
	result := make([][]rune, n)
	for i := 0; i < n; i++ {
		result[i] = make([]rune, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[i][j] = grid[j][n-1-i]
		}
	}

	return gridToString(result)
}

func SelfieDiag2Counterclock(s string) string {
	diag := Diag2Sym(s)
	rot := Rot90Counter(s)

	originalLines := strings.Split(s, "\n")
	diagLines := strings.Split(diag, "\n")
	rotLines := strings.Split(rot, "\n")

	result := make([]string, len(originalLines))
	for i := 0; i < len(originalLines); i++ {
		result[i] = originalLines[i] + "|" + diagLines[i] + "|" + rotLines[i]
	}

	return strings.Join(result, "\n")
}

type TransformFunc func(string) string

func Oper(fct TransformFunc, s string) string {
	return fct(s)
}

func stringToGrid(s string) [][]rune {
	lines := strings.Split(s, "\n")
	n := len(lines)
	grid := make([][]rune, n)
	for i := 0; i < n; i++ {
		grid[i] = []rune(lines[i])
	}
	return grid
}

func gridToString(grid [][]rune) string {
	lines := make([]string, len(grid))
	for i := 0; i < len(grid); i++ {
		lines[i] = string(grid[i])
	}
	return strings.Join(lines, "\n")
}