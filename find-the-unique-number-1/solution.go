package kata

func FindUniq(arr []float32) float32 {
	// compare first three to determine which value is the common one
	if arr[0] == arr[1] {
		for _, v := range arr[2:] {
			if v != arr[0] {
				return v
			}
		}
	} else {
		// arr[0] != arr[1], so one of them is unique; arr[2] breaks the tie
		if arr[0] == arr[2] {
			return arr[1]
		}
		return arr[0]
	}
	return arr[0]
}