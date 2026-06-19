// Let us define two sums v(n, p) and u(n, p):
//
//   v(n, p) = sum_{k=0}^{n} (-1)^k * p * 4^(n-k) * C(2n - k, k)
//   u(n, p) = sum_{k=0}^{n} (-1)^k * p * 4^(n-k) * C(2n - k + 1, k)
//
// where C(a, b) is the binomial coefficient.
//
// Task:
//  1) Calculate v(n, p) and u(n, p) with two brute-force functions V1(n, p) and U1(n, p).
//  2) Try V1(n, p) and U1(n, p) for small values of n and p and guess the closed forms.
//  3) Using 2) write VEff(n, p) and UEff(n, p) to efficiently calculate v and u
//     for bigger values of n and p.
//
// Examples:
//   V1(12, 70) --> 1750
//   U1(13, 18) --> 252
//
// See: https://en.wikipedia.org/wiki/Binomial_coefficient
package kata

import "math/big"

func V1(n, p *big.Int) *big.Int { return nil }

func U1(n, p *big.Int) *big.Int { return nil }

func VEff(n, p *big.Int) *big.Int { return nil }

func UEff(n, p *big.Int) *big.Int { return nil }