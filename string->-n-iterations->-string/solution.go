package main

func Repeat(s string, n int) string {
	seen := make(map[string]int)
	current := s
	seen[current] = 0
	
	for i := 1; i <= n; i++ {
		current = transform(current)
		
		if prevI, exists := seen[current]; exists {
			cycleLength := i - prevI
			remaining := (n - i) % cycleLength
			
			for j := 0; j < remaining; j++ {
				current = transform(current)
			}
			return current
		}
		
		seen[current] = i
	}
	
	return current
}

func transform(s string) string {
	runes := []rune(s)
	even := make([]rune, 0)
	odd := make([]rune, 0)
	
	for i, ch := range runes {
		if i%2 == 0 {
			even = append(even, ch)
		} else {
			odd = append(odd, ch)
		}
	}
	
	result := make([]rune, 0, len(runes))
	result = append(result, even...)
	result = append(result, odd...)
	
	return string(result)
}