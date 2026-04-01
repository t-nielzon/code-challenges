/*
You are given a string s made up of substring s(1), s(2), ..., s(n) separated by whitespaces.

Task: Return a string t having the following property:
length t(0) <= length t(1) >= length t(2) <= length t(3) >= length t(4) .... (P)

At each step from left to right, you can only move either already consecutive strings
or strings that became consecutive after a previous move. The number of moves should be minimum.

When you have property (P), t(0), t(2), ... (even indices) will be lower case
and t(1), t(3), ... (odd indices) will be upper case.

arrange("after be arrived two My so") should return "be ARRIVED two AFTER my SO"
*/
package kata

func Arrange(s string) string {
}