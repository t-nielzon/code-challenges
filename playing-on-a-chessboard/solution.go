package main

func game(n int) []interface{} {
	if n == 0 {
		return []interface{}{0}
	}

	numerator := int64(0)
	denominator := int64(1)

	for row := 1; row <= n; row++ {
		for col := 1; col <= n; col++ {
			fracNum := int64(col)
			fracDen := int64(col + row)

			numerator = numerator*fracDen + fracNum*denominator
			denominator = denominator * fracDen

			g := gcd(numerator, denominator)
			numerator /= g
			denominator /= g
		}
	}

	if denominator == 1 {
		return []interface{}{numerator}
	}
	return []interface{}{numerator, denominator}
}

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}