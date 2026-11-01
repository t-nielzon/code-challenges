// problem.go
/*
John has some amount of money of which he wants to deposit a part f0 to the bank at the beginning
of year 1. He wants to withdraw each year for his living an amount c0.

Banker's plan:
- deposit f0 at beginning of year 1
- bank account has an interest rate of p percent per year, constant over the years
- John withdraws each year c0, adjusted yearly for inflation rate i percent per year
- all amounts are truncated to their integral part (floor)
- determine if John can survive until the nth year

Given f0, p, c0, n, i, return true if John can make a living until the nth year, false otherwise.

Example:
fortune(100000, 1, 2000, 15, 1) -> true
fortune(100000, 1, 10000, 10, 1) -> true
fortune(100000, 1, 9185, 12, 1) -> false
*/

package main

func fortune(f0 int, p int, c0 int, n int, i int) bool {
	// stub
	return false
}