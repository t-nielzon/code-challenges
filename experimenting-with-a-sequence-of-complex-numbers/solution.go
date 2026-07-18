package main

import (
	"math"
	"math/cmplx"
)

func F(z complex128, eps float64) int {
	modZ := cmplx.Abs(z)
	
	// domain of convergence is |z| < 1
	if modZ >= 1 {
		return -1
	}
	
	// S(n, z) = z(1 - z^n), lim = z as n → ∞
	// condition: |z(1 - z^n) - z| < eps
	// simplifies to: |z|^(n+1) < eps
	// taking log: (n+1) * log|z| < log(eps)
	// since log|z| < 0: n+1 > log(eps) / log|z|
	// smallest integer n: floor(log(eps) / log|z|)
	
	logZ := math.Log(modZ)
	logEps := math.Log(eps)
	
	n := int(logEps / logZ)
	if n < 1 {
		n = 1
	}
	
	return n
}