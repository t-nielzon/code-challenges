package kata

func IsValidWalk(walk []byte) bool {
	if len(walk) != 10 {
		return false
	}
	x, y := 0, 0
	for _, d := range walk {
		switch d {
		case 'n':
			y++
		case 's':
			y--
		case 'e':
			x++
		case 'w':
			x--
		}
	}
	return x == 0 && y == 0
}