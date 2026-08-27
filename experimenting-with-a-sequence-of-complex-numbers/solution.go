package kata

import "math/cmplx"

func F(z complex128, eps float64) int {
	// Check if z is in the domain of convergence |z| < 1
	if cmplx.Abs(z) >= 1 {
		return -1
	}

	// The geometric series z + z^2 + ... + z^n = z(1 - z^n)/(1 - z)
	// So S(n, z) = (1 - z) * z(1 - z^n)/(1 - z) = z(1 - z^n)
	// As n → ∞ with |z| < 1: z^n → 0, so S(n, z) → z
	// We need the smallest n such that |S(n, z) - z| < eps
	// |z(1 - z^n) - z| = |z||−z^n| = |z|^(n+1) < eps
	
	zPowN := z  // z^1
	for n := 1; n <= 10000; n++ {
		// S(n, z) = z * (1 - z^n)
		sn := z * (1 - zPowN)

		// Check if |S(n, z) - z| < eps
		if cmplx.Abs(sn-z) < eps {
			return n
		}

		zPowN *= z  // z^(n+1)
	}

	return -1
}