package main

func FindUniq(arr []float64) float64 {
	// determine the common value by checking first 3 elements
	var common float64
	if arr[0] == arr[1] {
		common = arr[0]
	} else if arr[0] == arr[2] {
		common = arr[0]
	} else {
		common = arr[1]
	}
	
	// find and return the unique value
	for _, num := range arr {
		if num != common {
			return num
		}
	}
	
	return common
}