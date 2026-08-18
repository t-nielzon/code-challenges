package main

type Result struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(arr []int) Result {
	result := Result{
		Pos:   []int{},
		Peaks: []int{},
	}

	if len(arr) < 3 {
		return result
	}

	for i := 1; i < len(arr)-1; i++ {
		// Skip if current is less than next (definitely not a peak)
		if arr[i] < arr[i+1] {
			continue
		}

		// Find the end of the plateau
		j := i
		for j < len(arr)-1 && arr[j] == arr[j+1] {
			j++
		}

		// If plateau extends to the end, it's not a peak
		if j == len(arr)-1 {
			continue
		}

		// Check if this is a peak
		if arr[i-1] < arr[i] && arr[j] > arr[j+1] {
			result.Pos = append(result.Pos, i)
			result.Peaks = append(result.Peaks, arr[i])
		}

		i = j
	}

	return result
}