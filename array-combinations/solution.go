package kata

func Solve(arr [][]int) int {
	result := 1
	for _, sub := range arr {
		unique := make(map[int]bool)
		for _, v := range sub {
			unique[v] = true
		}
		result *= len(unique)
	}
	return result
}