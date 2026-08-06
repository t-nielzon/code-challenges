package main

/*
John has some amount of money of which he wants to deposit a part f0 to the bank at the beginning
of year 1. He wants to withdraw each year for his living an amount c0.

Here is his banker plan:
- deposit f0 at beginning of year 1
- his bank account has an interest rate of p percent per year, constant over the years
- John can withdraw each year c0, taking it whenever he wants in the year; he must take account of an inflation of i percent per year in order to keep his quality of living. i is supposed to stay constant over the years.
- all amounts f0..f(n-1), c0..c(n-1) are truncated by the bank to their integral part
- Given f0, p, c0, i the banker guarantees that John will be able to go on that way until the nth year.

John wants to know if the banker's plan is right or wrong.
Given parameters f0, p, c0, n, i build a function fortune which returns true if John can make a living until the nth year
and false if it is not possible.

Note: Don't forget to convert the percent parameters as percentages in the body of your function: if a parameter percent is 2 you have to convert it to 0.02.
*/

func fortune(f0 float64, p float64, c0 float64, n int, i float64) bool {
	return false
}