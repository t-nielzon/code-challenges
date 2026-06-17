package kata

func Thirt(n int) int {
	seq := []int{1, 10, 9, 12, 3, 4}
	for {
		sum := 0
		for i := 0; n > 0; i++ {
			sum += (n % 10) * seq[i%len(seq)]
			n /= 10
		}
		if sum == n {
			return sum
		}
		n = sum
	}
}