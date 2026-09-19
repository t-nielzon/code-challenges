package main

type PeakData struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(arr []int) PeakData {
	result := PeakData{
		Pos:   []int{},
		Peaks: []int{},
	}

	if len(arr) < 3 {
		return result
	}

	for i := 1; i < len(arr)-1; i++ {
		if arr[i] > arr[i-1] {
			// find end of plateau
			j := i
			for j+1 < len(arr) && arr[j] == arr[j+1] {
				j++
			}

			// check if value drops after plateau
			if j+1 < len(arr) && arr[j] > arr[j+1] {
				result.Pos = append(result.Pos, i)
				result.Peaks = append(result.Peaks, arr[i])
				i = j // skip past plateau
			}
		}
	}

	return result
}