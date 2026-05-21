package kata

func PickPeaks(arr []int) map[string][]int {
	pos := []int{}
	peaks := []int{}
	n := len(arr)
	for i := 1; i < n-1; i++ {
		if arr[i] > arr[i-1] {
			j := i
			for j < n-1 && arr[j+1] == arr[i] {
				j++
			}
			if j < n-1 && arr[j+1] < arr[i] {
				pos = append(pos, i)
				peaks = append(peaks, arr[i])
			}
			i = j
		}
	}
	return map[string][]int{"pos": pos, "peaks": peaks}
}