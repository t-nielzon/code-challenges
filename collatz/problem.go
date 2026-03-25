/*
Preface

A collatz sequence, starting with a positive integer n, is found by repeatedly
applying the following function to n until n == 1:

f(n) = n/2,  if n is even
f(n) = 3n+1, if n is odd

The Problem

Create a function collatz that returns a collatz sequence string starting with
the positive integer argument passed into the function, in the following form:

"X0->X1->...->XN"

Sample Input

Input: 4
Output: "4->2->1"

Input: 3
Output: "3->10->5->16->8->4->2->1"
*/
package kata

func Collatz(n int) string {
}