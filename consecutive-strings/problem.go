package main

/*
You are given an array(list) `strarr` of strings and an integer `k`.
Your task is to return the **first** longest string consisting of k **consecutive**
strings taken in the array.

Examples:
strarr = ["tree", "foling", "trashy", "blue", "abcdef", "uvwxyz"], k = 2

Concatenate the consecutive strings of strarr by 2:
treefoling   (length 10)  concatenation of strarr[0] and strarr[1]
folingtrashy (length 12)  concatenation of strarr[1] and strarr[2]
trashyblue   (length 10)  concatenation of strarr[2] and strarr[3]
blueabcdef   (length 10)  concatenation of strarr[3] and strarr[4]
abcdefuvwxyz (length 12) concatenation of strarr[4] and strarr[5]

Two strings are the longest: "folingtrashy" and "abcdefuvwxyz".
The first that came is "folingtrashy" so return "folingtrashy".

If n = 0 or k > n or k <= 0, return ""
*/

func LongestConsec(strarr []string, k int) string {
	// TODO: Implement the solution
	return ""
}