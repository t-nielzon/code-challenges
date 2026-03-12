package kata

func FindUniq(arr []float32) float32 {
	if arr[0] != arr[1] {
		if arr[0] == arr[2] {
			return arr[1]
		}
		return arr[0]
	}
	for i := 2; i < len(arr); i++ {
		if arr[i] != arr[0] {
			return arr[i]
		}
	}
	return arr[0]
}