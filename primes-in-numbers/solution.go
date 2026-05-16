package kata

import "fmt"

func PrimeFactors(n int) string {
	result := ""
	for p := 2; p*p <= n; p++ {
		count := 0
		for n%p == 0 {
			n /= p
			count++
		}
		if count == 1 {
			result += fmt.Sprintf("(%d)", p)
		} else if count > 1 {
			result += fmt.Sprintf("(%d**%d)", p, count)
		}
	}
	if n > 1 {
		result += fmt.Sprintf("(%d)", n)
	}
	return result
}