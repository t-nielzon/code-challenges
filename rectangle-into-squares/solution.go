package main

func SqInRect(lng, wdth int) []int {
	if lng == wdth {
		return nil
	}

	var result []int
	a, b := lng, wdth

	for b > 0 {
		times := a / b
		for i := 0; i < times; i++ {
			result = append(result, b)
		}
		a, b = b, a%b
	}

	return result
}