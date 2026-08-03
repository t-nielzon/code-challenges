package main

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int64) int64 {
	return a / gcd(a, b) * b
}

func ConvertFracs(fracs [][]int64) [][]int64 {
	if len(fracs) == 0 {
		return [][]int64{}
	}
	
	// Reduce input fractions first
	reduced := make([][]int64, len(fracs))
	for i, frac := range fracs {
		g := gcd(frac[0], frac[1])
		reduced[i] = []int64{frac[0] / g, frac[1] / g}
	}
	
	// Find LCM of all denominators
	var commonDenom int64 = reduced[0][1]
	for i := 1; i < len(reduced); i++ {
		commonDenom = lcm(commonDenom, reduced[i][1])
	}
	
	// Convert all fractions to have the common denominator
	result := make([][]int64, len(reduced))
	for i, frac := range reduced {
		num := frac[0] * (commonDenom / frac[1])
		result[i] = []int64{num, commonDenom}
	}
	
	return result
}