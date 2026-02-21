package kata

func SqInRect(lng, wdth int) []int {
	if lng == wdth {
		return nil
	}
	var result []int
	for lng != wdth {
		if lng > wdth {
			result = append(result, wdth)
			lng -= wdth
		} else {
			result = append(result, lng)
			wdth -= lng
		}
	}
	result = append(result, lng)
	return result
}