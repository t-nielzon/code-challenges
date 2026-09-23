package main

/*
Amidakuji is a method of lottery designed to create random pairings between two
sets comprised of an equal number of elements.

Your task is to write a function Amidakuji that returns the final positions of
each element. Note that the elements are an ascending sequence of consecutive
integers starting with 0 (from left to right).

Input: An array/list of equal-length strings consisting of 0 and 1 characters;
this represents the "ladder" structure. The 1s represent the rungs of the ladder
and the 0s represent empty space.

Each element begins at the top of its corresponding vertical rail. During the
descent of the ladder, whenever a vertical rail intersects a horizontal rung, it
swaps values with the adjacent connecting vertical rail.

Output: An array of integers, with each integer in its final position.

Example:
ladder := []string{
	"001001",
	"010000",
	"100100",
	"001000",
	"100101",
	"010010",
	"101001",
	"010100",
}

Amidakuji(ladder) // [4 2 0 5 3 6 1]
*/

func Amidakuji(ladder []string) []int {
	// stub
	return nil
}