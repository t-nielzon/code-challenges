/*
Given a string that contains only letters, you have to find out the number of
unique strings (including the string itself) that can be produced by
re-arranging the letters of the string. Strings are case insensitive.

Examples:
  UniqCount("AB") = 2      // "AB", "BA"
  UniqCount("ABC") = 6     // "ABC", "ACB", "BAC", "BCA", "CAB", "CBA"
  UniqCount("ABA") = 3     // "AAB", "ABA", "BAA"
  UniqCount("ABBb") = 4    // "ABBB", "BABB", "BBAB", "BBBA"
  UniqCount("AbcD") = 24   // "ABCD", etc.
*/
package kata

import "math/big"

func UniqCount(s string) *big.Int {
}