package main

import "sort"

func RemoveNb(n int) [][2]int {
	sumAll := int64(n) * int64(n+1) / 2
	target := sumAll + 1
	
	var result [][2]int
	
	// Find all divisors by iterating up to sqrt(target)
	for d := int64(1); d*d <= target; d++ {
		if target%d == 0 {
			// Check divisor d: (d-1, target/d - 1)
			a := d - 1
			b := target/d - 1
			if a >= 1 && a <= int64(n) && b >= 1 && b <= int64(n) && a != b {
				result = append(result, [2]int{int(a), int(b)})
			}
			
			// Check divisor target/d (if different from d): (target/d - 1, d - 1)
			if d != target/d {
				d2 := target / d
				a2 := d2 - 1
				b2 := d - 1
				if a2 >= 1 && a2 <= int64(n) && b2 >= 1 && b2 <= int64(n) && a2 != b2 {
					result = append(result, [2]int{int(a2), int(b2)})
				}
			}
		}
	}
	
	if len(result) == 0 {
		return nil
	}
	
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})
	
	return result
}