package main

/*
Common denominators

You will have a list of rationals in the form:
{numer_1, denom_1}, ..., {numer_n, denom_n}

All numbers are positive integers.

You have to produce a result where D is as small as possible and:
N_1/D == numer_1/denom_1 ... N_n/D == numer_n/denom_n.

Example:
ConvertFracs([][]int{{1, 2}, {1, 3}, {1, 4}}) should return [][]int{{6, 12}, {4, 12}, {3, 12}}

Note: Newer translations have reducible fractions. Simplify fractions to ensure D is minimal.
*/

func ConvertFracs(fracs [][]int) [][]int {
}