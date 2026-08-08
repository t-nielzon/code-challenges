package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func decompose(s string) []string {
	var num, den int64

	// Parse the input string
	if strings.Contains(s, "/") {
		// Fraction format: "n/d"
		parts := strings.Split(s, "/")
		if len(parts) != 2 {
			return []string{}
		}
		n, err1 := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
		d, err2 := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
		if err1 != nil || err2 != nil || d == 0 {
			return []string{}
		}
		num = n
		den = d
	} else if strings.Contains(s, ".") {
		// Decimal format: "0.6"
		parts := strings.Split(s, ".")
		if len(parts) != 2 {
			return []string{}
		}
		integer, err1 := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
		decimalStr := strings.TrimSpace(parts[1])
		decimalPart, err2 := strconv.ParseInt(decimalStr, 10, 64)
		if err1 != nil || err2 != nil {
			return []string{}
		}
		decimals := int64(len(decimalStr))
		den = int64(math.Pow10(int(decimals)))
		num = integer*den + decimalPart
	} else {
		// Integer format
		n, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
		if err != nil {
			return []string{}
		}
		num = n
		den = 1
	}

	if num == 0 {
		return []string{}
	}

	// GCD function for simplifying fractions
	gcd := func(a, b int64) int64 {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	// Simplify the fraction
	g := gcd(num, den)
	num /= g
	den /= g

	var result []string

	for num > 0 {
		if num >= den {
			// Add integer part
			intPart := num / den
			result = append(result, fmt.Sprintf("%d", intPart))
			num = num % den
		} else {
			// Find smallest d such that 1/d <= num/den
			// This means d >= den/num (ceiling division)
			d := (den + num - 1) / num
			result = append(result, fmt.Sprintf("1/%d", d))

			// Subtract 1/d from num/den: (num/den) - (1/d) = (num*d - den) / (den*d)
			num = num*d - den
			den = den * d

			// Simplify to prevent overflow
			g := gcd(num, den)
			num /= g
			den /= g
		}
	}

	return result
}

func main() {
	fmt.Println(decompose("21/23"))
	fmt.Println(decompose("2/3"))
	fmt.Println(decompose("0.6"))
}