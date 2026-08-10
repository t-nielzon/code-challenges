/*
The action of a Caesar cipher is to replace each plaintext letter (plaintext letters are from 'a' to 'z' or from 'A' to 'Z') with a different one a fixed number of places up or down the alphabet.

This program performs a variation of the Caesar shift. The shift increases by 1 for each character (on each iteration).

If the shift is initially 1, the first character of the message to be encoded will be shifted by 1, the second character will be shifted by 2, etc...

The function "movingShift" first codes the entire string and then returns an array of strings containing the coded string in 5 parts.

If possible the message will be equally divided between the five runners. If this is not possible, parts 1 to 4 will have non-increasing lengths such that parts 1 to 4 are at least as long as when evenly divided, but at most 1 longer.

The function "demovingShift" takes an array of strings and a shift value, and returns the decoded original string.

Example:
u = "I should have known that you would have a perfect answer for me!!!"
movingShift(u, 1) returns:
v = ["J vltasl rlhr ", "zdfog odxr ypw", " atasl rlhr p ", "gwkzzyq zntyhv", " lvz wp!!!"]
and demovingShift(v, 1) returns u.
*/

package main

func movingShift(s string, shift int) []string {
	// TODO: Implement
	return []string{}
}

func demovingShift(s []string, shift int) string {
	// TODO: Implement
	return ""
}