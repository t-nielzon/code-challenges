package main

import "strings"

// Moves in squared strings (IV)
//
// You are given a string of n lines, each substring being n characters long.
// We study transformations of this square of strings.
//
// - Symmetry with respect to the antidiagonal: diag_2_sym
// - Counterclockwise rotation 90 degrees: rot_90_counter
// - selfie_diag2_counterclock: concatenation of original, diag_2_sym, and rot_90_counter

func diag_2_sym(s string) string {
	// TODO
}

func rot_90_counter(s string) string {
	// TODO
}

func selfie_diag2_counterclock(s string) string {
	// TODO
}

func oper(fct func(string) string, s string) string {
	// TODO
}