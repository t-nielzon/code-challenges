/*
A Rule of Divisibility by 13

From Wikipedia:
"A divisibility rule is a shorthand way of determining whether a given
integer is divisible by a fixed divisor without performing the division,
usually by examining its digits."

When you divide the successive powers of 10 by 13 you get the following
remainders of the integer divisions:

1, 10, 9, 12, 3, 4

The cycle repeats. Multiply each digit (from the right) by the
corresponding value in this repeating sequence, sum the products,
and repeat until the result is stationary.

Task: Write a function which processes this sequence of operations on an
integer n (>=0) and returns the stationary number.
*/
package kata

func Thirt(n int) int {
}