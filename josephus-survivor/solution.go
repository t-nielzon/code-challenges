package kata

func JosSurvivor(n, k int) int {
	// iterative josephus formula: position is 0-indexed, convert to 1-indexed at the end
	survivor := 0
	for i := 2; i <= n; i++ {
		survivor = (survivor + k) % i
	}
	return survivor + 1
}