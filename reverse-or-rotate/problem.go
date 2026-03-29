/*
The input is a string str of digits. Cut the string into chunks of size sz
(ignore the last chunk if its size is less than sz).

If the sum of a chunk's digits is divisible by 2, reverse that chunk;
otherwise rotate it to the left by one position.
Put together these modified chunks and return the result as a string.

If sz is <= 0 or if str == "" return ""
If sz is greater than the length of str it is impossible to take a chunk
of size sz hence return "".
*/
package kata

func Revrot(s string, sz int) string {
}