package kata

// z^3 = x^2 * y^2 => k^3 = (x*y)^2 => x*y = k * sqrt(k)
// for integer solutions, k must be a perfect cube? no — k^3 must be a perfect square.
// k^3 is a perfect square iff all prime exponents in k are even after multiplying by 3,
// i.e., all prime exponents in k are even. that's too restrictive.
// actually k^3 = (x*y)^2, so (x*y) = k^(3/2). for x*y to be integer, k must be a perfect square.
// wait: reconsider. k^3 must be a perfect square. k = p1^a1 * p2^a2 * ...
// k^3 has exponents 3*ai. for perfect square, all 3*ai must be even, so all ai must be even.
// that means k must be a perfect square. let k = m^2, then k^3 = m^6 = (m^3)^2.
// so x*y = m^3. count positive integer pairs (x,y) with x*y = m^3, which is the number
// of divisors of m^3.
// but wait, k could also just need k^3 to be a perfect square. 3*ai even => ai even.
// so k = m^2 for some integer m, and we need divisor count of m^3.

func C(k int) int {
	// check if k is a perfect square
	m := isqrt(k)
	if m*m != k {
		return 0
	}

	// count divisors of m^3
	// factorize m, then for m^3 each exponent is tripled
	// number of divisors = product of (3*e + 1) for each prime factor
	return countDivisorsOfCube(m)
}

func isqrt(n int) int {
	if n <= 0 {
		return 0
	}
	x := int(0)
	// binary search
	lo, hi := 1, n
	if hi > 3162278 { // sqrt(10^13) ~ 3.16e6
		hi = 3162278
	}
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if mid <= n/mid { // mid*mid <= n, avoiding overflow
			x = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return x
}

func countDivisorsOfCube(m int) int {
	if m == 1 {
		return 1
	}
	result := 1
	n := m

	// trial division
	for p := 2; p*p <= n; p++ {
		if n%p == 0 {
			e := 0
			for n%p == 0 {
				e++
				n /= p
			}
			result *= (3*e + 1)
		}
	}
	if n > 1 {
		// remaining prime factor with exponent 1
		result *= (3*1 + 1)
	}
	return result
}