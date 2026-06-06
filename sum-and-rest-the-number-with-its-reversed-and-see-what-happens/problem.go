// The number 45 is the first integer in having this interesting property:
// the sum of the number with its reversed is divisible by the difference
// between them (absolute value).
//
//   45 + 54 = 99
//   abs(45 - 54) = 9
//   99 is divisible by 9.
//
// The first terms of this special sequence are:
//
//   n    a(n)
//   1    45
//   2    54
//   3    495
//   4    594
//
// Make the function that receives n, the ordinal number of the term, and
// gives us the value of the term of the sequence.
//
//   1 -----> 45
//   3 -----> 495
//
// Important: Do not include numbers which, when reversed, have a leading
// zero, e.g.: 1890 reversed is 0981, so 1890 should not be included.
//
// Your code will be tested up to the first 65 terms, so optimize it.

package kata

func SumDivSequence(n int) int {
	return 0
}