func ConvertFracs(fracs [][2]int) [][2]int {
	if len(fracs) == 0 {
		return [][2]int{}
	}

	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	lcm := func(a, b int) int {
		return (a / gcd(a, b)) * b
	}

	// reduce all fractions first
	reduced := make([][2]int, len(fracs))
	for i, frac := range fracs {
		g := gcd(frac[0], frac[1])
		reduced[i] = [2]int{frac[0] / g, frac[1] / g}
	}

	// find lcm of all denominators
	commonDenom := reduced[0][1]
	for i := 1; i < len(reduced); i++ {
		commonDenom = lcm(commonDenom, reduced[i][1])
	}

	// convert all fractions to common denominator
	result := make([][2]int, len(reduced))
	for i, frac := range reduced {
		newNum := (commonDenom / frac[1]) * frac[0]
		result[i] = [2]int{newNum, commonDenom}
	}

	return result
}