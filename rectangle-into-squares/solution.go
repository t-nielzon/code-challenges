func SqInRect(lng, wdth int) []int {
	if lng == wdth {
		return nil
	}

	var result []int
	l, w := lng, wdth

	for w > 0 {
		if l < w {
			l, w = w, l
		}
		count := l / w
		for i := 0; i < count; i++ {
			result = append(result, w)
		}
		l = l % w
	}

	return result
}