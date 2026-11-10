// Package main contains the solution for the Hamster Me kata.
//
// Hamsters are rodents belonging to the subfamily Cricetinae. The subfamily
// contains about 25 species, classified in six or seven genera. They have
// become established as popular small house pets.
//
// Write a function that accepts two inputs: code and message and returns an
// encrypted string from message using the code.
//
// The code is a string that generates a cipher key where:
// - All letters from code get level 1
// - All letters which directly follow letters from code get level 2 (unless
//   they already have a smaller number assigned), etc.
//
// Example cipher for code "hamster":
//  1  | h a m s t e r
//  2  | i b n   u f
//  3  | j c o   v g
//  4  | k d p   w
//  5  | l   q   x
//  6  |         y
//  7  |         z
//
// The encoding works like:
//  a => a1, b => a2, c => a3, e => e1, f => e2, etc.
//
// Examples:
//  hamsterMe('hamster', 'hamster') => h1a1m1s1t1e1r1
//  hamsterMe('hamster', 'helpme')  => h1e1h5m4m1e1
//
// If a letter is not in the code, find the largest code letter that is <= it,
// then calculate the distance. If no such letter exists, wrap around to the
// largest code letter.
package main

func HamsterMe(code, message string) string {
	// TODO: implement
	return ""
}