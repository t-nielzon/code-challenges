/*
A Rule of Divisibility by 13

From Wikipedia:
"A divisibility rule is a shorthand way of determining whether a given integer is divisible by a fixed divisor without performing the division, usually by examining its digits."

When you divide the successive powers of 10 by 13 you get the following remainders of the integer divisions:
1, 10, 9, 12, 3, 4 because:
10^0 ->  1 (mod 13)
10^1 -> 10 (mod 13)
10^2 ->  9 (mod 13)
10^3 -> 12 (mod 13)
10^4 ->  3 (mod 13)
10^5 ->  4 (mod 13)

Then the whole pattern repeats.

Multiply the rightmost digit of the number with the leftmost number in the sequence,
the second rightmost digit with the second leftmost digit of the sequence, and so on.
The cycle repeats, and you sum all these products. Repeat this process until the sequence
of sums becomes stationary (doesn't change anymore).

Example with 1234567:
7×1 + 6×10 + 5×9 + 4×12 + 3×3 + 2×4 + 1×1 = 178
8×1 + 7×10 + 1×9 = 87
7×1 + 8×10 = 87 (stationary)

Write a function which processes this sequence of operations on an integer n (>=0)
and returns the stationary number.
*/

package main

func DivisibleBy13(n int) int {
}