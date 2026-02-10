package kata

func BouncingBall(h, bounce, window float64) int {
	if h <= 0 || bounce <= 0 || bounce >= 1 || window >= h {
		return -1
	}

	count := 1 // the initial drop passes the window
	h *= bounce

	for h > window {
		count += 2 // once going up, once coming down
		h *= bounce
	}

	return count
}