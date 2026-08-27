package main

func Transform(s string, n int) string {
	if n == 0 {
		return s
	}
	
	seen := make(map[string]int)
	current := s
	iteration := 0
	seen[current] = iteration
	
	for iteration < n {
		current = transformOnce(current)
		iteration++
		
		if prevIteration, ok := seen[current]; ok {
			cycleLength := iteration - prevIteration
			offset := (n - prevIteration) % cycleLength
			
			result := current
			for i := 0; i < offset; i++ {
				result = transformOnce(result)
			}
			return result
		}
		
		seen[current] = iteration
	}
	
	return current
}

func transformOnce(s string) string {
	even := []rune{}
	odd := []rune{}
	
	for i, ch := range s {
		if i%2 == 0 {
			even = append(even, ch)
		} else {
			odd = append(odd, ch)
		}
	}
	
	return string(append(even, odd...))
}