package kata

import "fmt"

func CuckooClock(initialTime string, n int) string {
	var h, m int
	fmt.Sscanf(initialTime, "%d:%d", &h, &m)
	for {
		chimes := 0
		switch m {
		case 0:
			chimes = h
		case 15, 30, 45:
			chimes = 1
		}
		if chimes > 0 {
			if n <= chimes {
				return fmt.Sprintf("%02d:%02d", h, m)
			}
			n -= chimes
		}
		switch {
		case m < 15:
			m = 15
		case m < 30:
			m = 30
		case m < 45:
			m = 45
		default:
			m = 0
			h++
			if h > 12 {
				h = 1
			}
		}
	}
}