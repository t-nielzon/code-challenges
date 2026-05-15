package kata

import "math"

func GiveTriang(p int) int {
	count := 0
	for a := 1; a < p; a++ {
		for b := a; a+b < p; b++ {
			c2 := a*a + b*b + a*b
			c := int(math.Sqrt(float64(c2)))
			for _, cc := range []int{c - 1, c, c + 1} {
				if cc > 0 && cc*cc == c2 && a+b+cc <= p {
					count++
					break
				}
			}
		}
	}
	return count
}