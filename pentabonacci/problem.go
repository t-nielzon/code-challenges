/*
We have the following sequence:

f(0) = 0
f(1) = 1
f(2) = 1
f(3) = 2
f(4) = 4
f(n) = f(n-1) + f(n-2) + f(n-3) + f(n-4) + f(n-5)

Your task is to give the number of total values for the odd terms of the
sequence up to the n-th term (included). (The number n (of n-th term) will
be given as a nonnegative integer)

The values 1 (one) is the only that is duplicated in the sequence and should
be counted only once.

E.g.
count_odd_pentaFib(5) -----> 1
count_odd_pentaFib(10) ------> 3
count_odd_pentaFib(15) ------> 5
*/
package kata

func CountOddPentaFib(n int) int {
}