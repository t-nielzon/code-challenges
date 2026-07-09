/*
Given a string that includes alphanumeric characters ("3a4B2d") return the 
expansion of that string: The numeric values represent the occurrence of each 
letter following that numeric value. There should be no numeric characters in 
the final string.

Notes:
- The first occurrence of a numeric value should be the number of times each 
  character behind it is repeated, until the next numeric value appears
- If there are multiple consecutive numeric characters, only the last one should 
  be used (ignore the previous ones)
- Empty strings should return an empty string

Examples:
"3D2a5d2f"  -->  "DDDaadddddff"    # 3*D + 2*a + 5*d + 2*f
"3abc"      -->  "aaabbbccc"       # 3*a + 3*b + 3*c
"3d332f2a"  -->  "dddffaa"         # 3*d + 2*f + 2*a (33 becomes 3)
"abcde"     -->  "abcde"           # no digits
"1111"      -->  ""                # no characters to repeat
""          -->  ""                # empty string
*/

function simpleStringExpansion(str) {
  // TODO: implement
}