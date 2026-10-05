// solution.go
func Amidakuji(ladder []string) []int {
	n := len(ladder[0])
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}

	for _, row := range ladder {
		for i := 0; i < n-1; i++ {
			if row[i] == '1' {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}

	return arr
}