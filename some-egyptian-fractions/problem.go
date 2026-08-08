package main

/*
Given a rational number n (n >= 0, with denominator strictly positive) as a string
(e.g., "2/3") decompose this number as a sum of rationals with numerators equal to one
and without repetitions.

The algorithm must be "greedy", so at each stage the new rational obtained in the
decomposition must have a denominator as small as possible.

Example:
decompose("21/23") should return ["1/2", "1/3", "1/13", "1/359", "1/644046"]

Notes:
1) The rational given could be greater than or equal to one, in which case the first
   "fraction" will be an integer (with implicit denominator of 1).
2) If the numerator parses to zero the function returns [] (empty slice).
3) The number could also be a decimal like "0.6".
*/

func decompose(s string) []string {
	return []string{}
}