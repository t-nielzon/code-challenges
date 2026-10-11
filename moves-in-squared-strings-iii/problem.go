package main

import "strings"

// You are given a string of n lines, each substring being n characters long.
// We will study some transformations of this square of strings.
//
// Symmetry with respect to the main diagonal: diag1Sym
// diag1Sym(s) => "aeim\nbfjn\ncgko\ndhlp"
//
// Clockwise rotation 90 degrees: rot90Clock
// rot90Clock(s) => "miea\nnjfb\nokgc\nplhd"
//
// selfieAndDiag1(s)
// It is initial string + string obtained by symmetry with respect to the main diagonal.
// selfieAndDiag1(s) => "abcd|aeim\nefgh|bfjn\nijkl|cgko\nmnop|dhlp"
//
// Task: Write these functions and a high-order function oper(fct, s) where
// fct is the function to apply to the string s.

func diag1Sym(s string) string {
	// stub
}

func rot90Clock(s string) string {
	// stub
}

func selfieAndDiag1(s string) string {
	// stub
}

func oper(fct func(string) string, s string) string {
	// stub
}