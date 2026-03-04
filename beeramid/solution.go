package kata

func Beeramid(bonus int, price float64) int {
	if bonus <= 0 || price <= 0 {
		return 0
	}
	budget := float64(bonus)
	level := 0
	spent := 0.0
	for {
		level++
		spent += float64(level*level) * price
		if spent > budget {
			return level - 1
		}
	}
}