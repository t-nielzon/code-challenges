package kata

func RemoveNb(n int) [][2]int {
	sum := n * (n + 1) / 2
	var result [][2]int
	for a := 1; a <= n; a++ {
		// a*b = sum - a - b  =>  b = (sum - a) / (a + 1)
		if (sum-a)%(a+1) == 0 {
			b := (sum - a) / (a + 1)
			if b >= 1 && b <= n {
				result = append(result, [2]int{a, b})
			}
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}