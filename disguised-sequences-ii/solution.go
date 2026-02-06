package kata

import "math/big"

// binomial computes C(n, k) using big.Int
func binomial(n, k int) *big.Int {
	if k < 0 || k > n {
		return big.NewInt(0)
	}
	if k == 0 || k == n {
		return big.NewInt(1)
	}
	// use symmetry to reduce iterations
	if k > n-k {
		k = n - k
	}
	result := big.NewInt(1)
	for i := 0; i < k; i++ {
		result.Mul(result, big.NewInt(int64(n-i)))
		result.Div(result, big.NewInt(int64(i+1)))
	}
	return result
}

// V1 computes v(n,p) using brute-force summation
func V1(n, p int) *big.Int {
	result := big.NewInt(0)
	four := big.NewInt(4)
	
	for k := 0; k <= n; k++ {
		// (-1)^k * p * 4^(n-k) * C(2n-k, k)
		term := big.NewInt(int64(p))
		
		// 4^(n-k)
		power := new(big.Int).Exp(four, big.NewInt(int64(n-k)), nil)
		term.Mul(term, power)
		
		// C(2n-k, k)
		binom := binomial(2*n-k, k)
		term.Mul(term, binom)
		
		// (-1)^k
		if k%2 == 1 {
			term.Neg(term)
		}
		
		result.Add(result, term)
	}
	return result
}

// U1 computes u(n,p) using brute-force summation
func U1(n, p int) *big.Int {
	result := big.NewInt(0)
	four := big.NewInt(4)
	
	for k := 0; k <= n; k++ {
		// (-1)^k * p * 4^(n-k) * C(2n-k+1, k)
		term := big.NewInt(int64(p))
		
		// 4^(n-k)
		power := new(big.Int).Exp(four, big.NewInt(int64(n-k)), nil)
		term.Mul(term, power)
		
		// C(2n-k+1, k)
		binom := binomial(2*n-k+1, k)
		term.Mul(term, binom)
		
		// (-1)^k
		if k%2 == 1 {
			term.Neg(term)
		}
		
		result.Add(result, term)
	}
	return result
}

// VEff computes v(n,p) efficiently
// pattern discovered: v(n, p) = p * (n + 1)
func VEff(n, p int) *big.Int {
	result := big.NewInt(int64(n + 1))
	result.Mul(result, big.NewInt(int64(p)))
	return result
}

// UEff computes u(n,p) efficiently
// pattern discovered: u(n, p) = p * (n + 1)
func UEff(n, p int) *big.Int {
	result := big.NewInt(int64(n + 1))
	result.Mul(result, big.NewInt(int64(p)))
	return result
}