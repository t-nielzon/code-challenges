/*
The numbers 12, 63 and 119 have something in common related with their divisors
and their prime factors, let's see it.

Numbers PrimeFactorsSum(pfs)        DivisorsSum(ds)              Is ds divisible by pfs
12         2 + 2 + 3 = 7         1 + 2 + 3 + 4 + 6 + 12 = 28            28 / 7 = 4,  Yes
63         3 + 3 + 7 = 13        1 + 3 + 7 + 9 + 21 + 63 = 104         104 / 13 = 8, Yes
119        7 + 17 = 24           1 + 7 + 17 + 119 = 144                144 / 24 = 6, Yes

There is an obvious property: the sum of the divisors of a number is divisible
by the sum of its prime factors.

ds_multof_pfs(nMin, nMax) returns a sorted list of numbers in [nMin, nMax]
where the sum of divisors is divisible by the sum of prime factors.

ds_multof_pfs(10, 100) == [12, 15, 35, 42, 60, 63, 66, 68, 84, 90, 95]
ds_multof_pfs(20, 120) == [35, 42, 60, 63, 66, 68, 84, 90, 95, 110, 114, 119]
*/
package kata

func DsMultofPfs(nMin, nMax int) []int {
}