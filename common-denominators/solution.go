package main

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func ConvertFracs(fracs [][]int) [][]int {
	if len(fracs) == 0 {
		return [][]int{}
	}

	// Simplify all input fractions
	simplified := make([][]int, len(fracs))
	for i, frac := range fracs {
		g := gcd(frac[0], frac[1])
		simplified[i] = []int{frac[0] / g, frac[1] / g}
	}

	// Find LCM of all simplified denominators
	commonDenom := simplified[0][1]
	for i := 1; i < len(simplified); i++ {
		commonDenom = lcm(commonDenom, simplified[i][1])
	}

	// Convert all fractions to have the common denominator
	result := make([][]int, len(simplified))
	for i, frac := range simplified {
		newNumer := frac[0] * (commonDenom / frac[1])
		result[i] = []int{newNumer, commonDenom}
	}

	return result
}