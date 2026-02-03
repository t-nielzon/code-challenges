package kata

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(array []int) PosPeaks {
	result := PosPeaks{Pos: []int{}, Peaks: []int{}}
	
	if len(array) < 3 {
		return result
	}
	
	peakStart := -1
	
	for i := 1; i < len(array)-1; i++ {
		if array[i] > array[i-1] {
			// Start of a potential peak or plateau
			peakStart = i
		} else if array[i] < array[i-1] {
			// End of a potential peak
			if peakStart != -1 {
				result.Pos = append(result.Pos, peakStart)
				result.Peaks = append(result.Peaks, array[peakStart])
				peakStart = -1
			}
		}
		// If array[i] == array[i-1], we're on a plateau, keep peakStart unchanged
	}
	
	return result
}