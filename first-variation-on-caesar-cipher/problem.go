// The action of a Caesar cipher is to replace each plaintext letter
// (plaintext letters are from 'a' to 'z' or from 'A' to 'Z') with a
// different one a fixed number of places up or down the alphabet.
//
// This program performs a variation of the Caesar shift. The shift
// increases by 1 for each character (on each iteration).
//
// If the shift is initially 1, the first character of the message to be
// encoded will be shifted by 1, the second character will be shifted by 2,
// etc...
//
// movingShift(s, shift):
//   param s: a string to be coded
//   param shift: an integer giving the initial shift
//   First codes the entire string, then returns an array of strings
//   containing the coded string in 5 parts. If possible the message is
//   equally divided by length between the five runners. If not possible,
//   parts 1 to 5 have non-increasing lengths, parts 1 to 4 at least as long
//   as when evenly divided but at most 1 longer. An empty final part must
//   still be shown.
//
//   Example: u = "I should have known that you would have a perfect answer for me!!!"
//   movingShift(u, 1) returns:
//   ["J vltasl rlhr ", "zdfog odxr ypw", " atasl rlhr p ", "gwkzzyq zntyhv", " lvz wp!!!"]
//
// demovingShift(s, shift):
//   param s: an array of strings (5 strings, possibly from movingShift)
//   param shift: an int
//   returns the decoded string.
//
// demovingShift(v, 1) returns u.

package kata

func movingShift(s string, shift int) []string {
	return nil
}

func demovingShift(s []string, shift int) string {
	return ""
}