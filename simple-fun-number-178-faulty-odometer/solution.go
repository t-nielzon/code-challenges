package kata

import "strconv"

// the odometer skips digit 4 in every position, so each digit is effectively
// a base-9 symbol drawn from {0,1,2,3,5,6,7,8,9}. interpret the displayed
// number as a base-9 value to recover the real distance.
func FaultyOdometer(n int) int {
	s := strconv.Itoa(n)
	result := 0
	for _, c := range s {
		d := int(c - '0')
		// digits above the skipped 4 are shifted down by one
		if d > 4 {
			d--
		}
		result = result*9 + d
	}
	return result
}