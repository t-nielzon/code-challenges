package main

func UnToZeroOrInfinity(n int) float64 {
	// Simplify: un = (1/n!) * (1! + 2! + ... + n!)
	//         = 1!/n! + 2!/n! + ... + n!/n!
	//         = 1/((2)(3)...(n)) + 1/((3)(4)...(n)) + ... + 1/n + 1
	//
	// Compute iteratively from k=n down to k=1:
	// term_k = 1/((k+1)(k+2)...(n))
	// Each term_k = term_{k+1} / (k+1)
	
	result := 1.0
	term := 1.0
	
	for k := n - 1; k >= 1; k-- {
		term /= float64(k + 1)
		result += term
	}
	
	return result
}