package kata

func Decode(roman string) int {
	values := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	for i := 0; i < len(roman); i++ {
		curr := values[roman[i]]
		if i+1 < len(roman) && curr < values[roman[i+1]] {
			total -= curr
		} else {
			total += curr
		}
	}
	return total
}