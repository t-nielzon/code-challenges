package kata

func sqInRect(lng int, wdth int) []int {
	if lng == wdth {
		return nil
	}
	var res []int
	for lng > 0 && wdth > 0 {
		if lng > wdth {
			res = append(res, wdth)
			lng -= wdth
		} else {
			res = append(res, lng)
			wdth -= lng
		}
	}
	return res
}