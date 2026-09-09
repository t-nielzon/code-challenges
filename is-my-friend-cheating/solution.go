package main

func RemoveNb(n int64) [][2]int64 {
	// (a+1)(b+1) = n(n+1)/2 + 1
	target := n*(n+1)/2 + 1
	
	var result [][2]int64
	
	// find all divisor pairs
	for d1 := int64(1); d1*d1 < target; d1++ {
		if target%d1 == 0 {
			d2 := target / d1
			a := d1 - 1
			b := d2 - 1
			
			// check if both are in valid range [1, n]
			if a >= 1 && a <= n && b >= 1 && b <= n {
				result = append(result, [2]int64{a, b})
				result = append(result, [2]int64{b, a})
			}
		}
	}
	
	return result
}