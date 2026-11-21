package main

func sqInRect(lng, wdth int) []int {
	if lng == wdth {
		return nil
	}

	var result []int
	for lng > 0 && wdth > 0 {
		if lng == wdth {
			result = append(result, lng)
			break
		}

		if lng > wdth {
			result = append(result, wdth)
			lng -= wdth
		} else {
			result = append(result, lng)
			wdth -= lng
		}
	}

	return result
}