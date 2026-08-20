package main

func DeadAnts(s string) int {
	count_a := 0
	count_n := 0
	count_t := 0

	for _, ch := range s {
		switch ch {
		case 'a':
			count_a++
		case 'n':
			count_n++
		case 't':
			count_t++
		}
	}

	// maximum possible ants is determined by the letter that appears least
	maxAnts := count_a
	if count_n < maxAnts {
		maxAnts = count_n
	}
	if count_t < maxAnts {
		maxAnts = count_t
	}

	// count living ants (non-overlapping "ant" substrings)
	livingAnts := 0
	for i := 0; i < len(s)-2; i++ {
		if s[i:i+3] == "ant" {
			livingAnts++
			i += 2 // skip to avoid counting overlapping occurrences
		}
	}

	return maxAnts - livingAnts
}