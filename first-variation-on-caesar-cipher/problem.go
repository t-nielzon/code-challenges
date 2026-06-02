// The action of a Caesar cipher is to replace each plaintext letter (plaintext
// letters are from 'a' to 'z' or from 'A' to 'Z') with a different one a fixed
// number of places up or down the alphabet.
//
// This program performs a variation of the Caesar shift. The shift increases
// by 1 for each character (on each iteration).
//
// If the shift is initially 1, the first character of the message to be encoded
// will be shifted by 1, the second character will be shifted by 2, etc...
//
// movingShift(s string, shift int) first codes the entire string and then
// returns an array of 5 strings containing the coded string in 5 parts.
// The message is split into 5 parts of size ceil(len/5) for the first parts,
// with the remainder going to the later parts (possibly empty).
//
// demovingShift(s []string, shift int) takes the 5-part array and the initial
// shift and returns the decoded string.
//
// Example:
//   u = "I should have known that you would have a perfect answer for me!!!"
//   movingShift(u, 1) ->
//     ["J vltasl rlhr ", "zdfog odxr ypw", " atasl rlhr p ", "gwkzzyq zntyhv", " lvz wp!!!"]
//   demovingShift(movingShift(u, 1), 1) -> u

package kata

func movingShift(s string, shift int) []string {
	// stub
	return nil
}

func demovingShift(s []string, shift int) string {
	// stub
	return ""
}