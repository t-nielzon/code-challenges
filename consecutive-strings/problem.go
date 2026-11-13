package main

/*
Consecutive strings

You are given an array(list) of strings and an integer k.
Your task is to return the first longest string consisting of k consecutive strings taken in the array.

Examples:
strarr = ["tree", "foling", "trashy", "blue", "abcdef", "uvwxyz"], k = 2

Concatenate consecutive strings of strarr by 2:
- treefoling (length 10)
- folingtrashy (length 12)
- trashyblue (length 10)
- blueabcdef (length 10)
- abcdefuvwxyz (length 12)

Two strings are longest: "folingtrashy" and "abcdefuvwxyz".
The first is "folingtrashy", so longest_consec(strarr, 2) returns "folingtrashy".

Rules:
- n being the length of the string array
- if n = 0 or k > n or k <= 0 return ""
*/

func LongestConsec(strarr []string, k int) string {
	// TODO: implement
	return ""
}