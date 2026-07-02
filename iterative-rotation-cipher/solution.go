package solution

import (
	"strconv"
	"strings"
)

func rotateRight(s []rune, n int) []rune {
	if len(s) == 0 {
		return s
	}
	k := n % len(s)
	res := make([]rune, 0, len(s))
	res = append(res, s[len(s)-k:]...)
	res = append(res, s[:len(s)-k]...)
	return res
}

func rotateLeft(s []rune, n int) []rune {
	if len(s) == 0 {
		return s
	}
	k := n % len(s)
	res := make([]rune, 0, len(s))
	res = append(res, s[k:]...)
	res = append(res, s[:k]...)
	return res
}

// reinsert space runes at the recorded positions of the full-length string
func putSpaces(noSpace []rune, positions []int) []rune {
	total := len(noSpace) + len(positions)
	res := make([]rune, 0, total)
	pi, ni := 0, 0
	for i := 0; i < total; i++ {
		if pi < len(positions) && positions[pi] == i {
			res = append(res, ' ')
			pi++
		} else {
			res = append(res, noSpace[ni])
			ni++
		}
	}
	return res
}

// apply a rotation to each maximal run of non-space characters
func rotateSubstrings(r []rune, n int, rot func([]rune, int) []rune) []rune {
	res := make([]rune, 0, len(r))
	i := 0
	for i < len(r) {
		if r[i] == ' ' {
			res = append(res, ' ')
			i++
			continue
		}
		j := i
		for j < len(r) && r[j] != ' ' {
			j++
		}
		res = append(res, rot(r[i:j], n)...)
		i = j
	}
	return res
}

func Encode(n int, s string) string {
	r := []rune(s)
	for iter := 0; iter < n; iter++ {
		var noSpace []rune
		var positions []int
		for i, c := range r {
			if c == ' ' {
				positions = append(positions, i)
			} else {
				noSpace = append(noSpace, c)
			}
		}
		noSpace = rotateRight(noSpace, n)
		r = putSpaces(noSpace, positions)
		r = rotateSubstrings(r, n, rotateRight)
	}
	return strconv.Itoa(n) + " " + string(r)
}

func Decode(s string) string {
	idx := strings.Index(s, " ")
	n, _ := strconv.Atoi(s[:idx])
	r := []rune(s[idx+1:])
	for iter := 0; iter < n; iter++ {
		r = rotateSubstrings(r, n, rotateLeft)
		var noSpace []rune
		var positions []int
		for i, c := range r {
			if c == ' ' {
				positions = append(positions, i)
			} else {
				noSpace = append(noSpace, c)
			}
		}
		noSpace = rotateLeft(noSpace, n)
		r = putSpaces(noSpace, positions)
	}
	return string(r)
}