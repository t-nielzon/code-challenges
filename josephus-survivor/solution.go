package kata

func JosephusSurvivor(n, k int) int {
	survivor := 0
	for i := 2; i <= n; i++ {
		survivor = (survivor + k) % i
	}
	return survivor + 1
}