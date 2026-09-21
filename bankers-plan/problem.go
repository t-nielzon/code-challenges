/*
Banker's Plan

John has some amount of money of which he wants to deposit a part f0 to the bank at the beginning
of year 1. He wants to withdraw each year for his living an amount c0.

His banker plan:
- deposit f0 at beginning of year 1
- his bank account has an interest rate of p percent per year
- John can withdraw each year c0, adjusted for inflation of i percent per year
- all amounts f0..f(n-1), c0..c(n-1) are truncated to their integral part
- the banker guarantees John can go on until the nth year

Given f0, p, c0, n, i, return true if John can make a living until the nth year, false otherwise.
Note: Convert percent parameters to decimal (2 becomes 0.02).
*/
package main

func fortune(f0 float64, p float64, c0 float64, n int, i float64) bool {
	return false
}