/*
Let us define two sums v(n, p) and u(n, p):

v(n,p) = Σ(k=0 to n) (-1)^k × p × 4^(n-k) × C(2n-k, k)

u(n,p) = Σ(k=0 to n) (-1)^k × p × 4^(n-k) × C(2n-k+1, k)

Task:
1) Calculate v(n, p) and u(n, p) with two brute-force functions v1(n, p) and u1(n, p).
2) Try v1(n, p) and u1(n, p) for small values of n and p and guess the results of u(n, p) and v(n, p)
3) Using 2) write VEff(n, p) and UEff(n, p) to efficiently calculate v and u for bigger values of n and p.

Examples:
V1(12, 70) --> 1750
U1(13, 18) --> 252
*/
package kata

import "math/big"

func V1(n, p int) *big.Int {
}

func U1(n, p int) *big.Int {
}

func VEff(n, p int) *big.Int {
}

func UEff(n, p int) *big.Int {
}