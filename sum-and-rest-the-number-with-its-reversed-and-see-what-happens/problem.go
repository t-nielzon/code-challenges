/*
Sum and Rest the Number with its Reversed and See What Happens

The number 45 is the first integer in having this interesting property:
the sum of the number with its reversed is divisible by the difference
between them (absolute value).

45 + 54 = 99
abs(45 - 54) = 9
99 is divisible by 9.

The first terms of this special sequence are:
n       a(n)
1       45
2       54
3       495
4       594

Make the function that receives n, the ordinal number of the term and
may give us, the value of the term of the sequence.

Important: Do not include numbers which, when reversed, have a leading
zero, e.g.: 1890 reversed is 0981, so 1890 should not be included.
*/
package kata

func SumAndRest(n int) int64 {
}