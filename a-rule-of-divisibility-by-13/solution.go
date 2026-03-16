package kata

func Thirt(n int) int {
	seq := []int{1, 10, 9, 12, 3, 4}
	prev := -1
	cur := n
	for cur != prev {
		prev = cur
		sum := 0
		tmp := cur
		i := 0
		for tmp > 0 {
			sum += (tmp % 10) * seq[i%6]
			tmp /= 10
			i++
		}
		cur = sum
	}
	return cur
}