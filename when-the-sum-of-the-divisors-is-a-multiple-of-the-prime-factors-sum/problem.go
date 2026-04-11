// problem.go
package kata

/*
The numbers 12, 63 and 119 have something in common related with their
divisors and their prime factors, let's see it.

Numbers PrimeFactorsSum(pfs)        DivisorsSum(ds)              Is ds divisible by pfs
12         2 + 2 + 3 = 7         1 + 2 + 3 + 4 + 6 + 12 = 28            28 / 7 = 4,  Yes
63         3 + 3 + 7 = 13        1 + 3 + 7 + 9 + 21 + 63 = 104         104 / 13 = 8, Yes
119        7 + 17 = 24           1 + 7 + 17 + 119 = 144                144 / 24 = 6, Yes

There is an obvious property: the sum of the divisors of a number is
divisible by the sum of its prime factors.

ds_multof_pfs(nMin, nMax) returns a sorted list of numbers in [nMin, nMax]
that fulfill this property.
*/

func DsMultofPfs(nMin, nMax int) []int {
}