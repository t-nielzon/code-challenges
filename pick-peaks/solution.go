package solution

func PickPeaks(arr []int) map[string][]int {
	result := map[string][]int{
		"pos":   []int{},
		"peaks": []int{},
	}

	if len(arr) < 3 {
		return result
	}

	for i := 1; i < len(arr)-1; i++ {
		if arr[i] > arr[i-1] {
			// find end of plateau
			j := i
			for j < len(arr)-1 && arr[j] == arr[j+1] {
				j++
			}

			// peak if followed by smaller value
			if j < len(arr)-1 && arr[j] > arr[j+1] {
				result["pos"] = append(result["pos"], i)
				result["peaks"] = append(result["peaks"], arr[i])
				i = j // skip past plateau
			}
		}
	}

	return result
}