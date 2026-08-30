package main

// Consider the following numbers (where n! is factorial(n)):
// u1 = (1 / 1!) * (1!)
// u2 = (1 / 2!) * (1! + 2!)
// u3 = (1 / 3!) * (1! + 2! + 3!)
// ...
// un = (1 / n!) * (1! + 2! + 3! + ... + n!)
//
// Calculate (1 / n!) * (1! + 2! + 3! + ... + n!) for a given n.
// The result should be within 10^-6 of the expected value.
//
// Factorials grow rapidly, so the expression needs to be simplified
// to handle large inputs without overflow.

func UnToZeroOrInfinity(n int) float64 {
}