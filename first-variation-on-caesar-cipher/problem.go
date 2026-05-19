package kata

/*
The action of a Caesar cipher is to replace each plaintext letter (plaintext letters are from 'a' to 'z' or from 'A' to 'Z') with a different one a fixed number of places up or down the alphabet.

This program performs a variation of the Caesar shift. The shift increases by 1 for each character (on each iteration).

If the shift is initially 1, the first character of the message to be encoded will be shifted by 1, the second character will be shifted by 2, etc...

Coding: Parameters and return of function "movingShift"
  param s: a string to be coded
  param shift: an integer giving the initial shift

The function "movingShift" first codes the entire string and then returns an array of strings containing the coded string in 5 parts (five parts because, to avoid more risks, the coded message will be given to five runners, one piece for each runner).

If possible the message will be equally divided by message length between the five runners. If this is not possible, parts 1 to 5 will have subsequently non-increasing lengths, such that parts 1 to 4 are at least as long as when evenly divided, but at most 1 longer. If the last part is the empty string this empty string must be shown in the resulting array.

For example, if the coded message has a length of 17 the five parts will have lengths of 4, 4, 4, 4, 1. If the length is 16 the parts will be of lengths 4, 4, 4, 4, 0. If the length is 11, parts will be of lengths 3, 3, 3, 2, 0.

Decoding: parameters and return of function "demovingShift"
  1) an array of strings: s (possibly resulting from "movingShift", with 5 strings)
  2) an int shift
"demovingShift" returns a string.

Example:
u = "I should have known that you would have a perfect answer for me!!!"
movingShift(u, 1) returns:
v = ["J vltasl rlhr ", "zdfog odxr ypw", " atasl rlhr p ", "gwkzzyq zntyhv", " lvz wp!!!"]
and demovingShift(v, 1) returns u.
*/

func MovingShift(s string, shift int) [5]string {
}

func DemovingShift(s [5]string, shift int) string {
}