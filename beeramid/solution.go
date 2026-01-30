package kata

func Beeramid(bonus int, price float64) int {
	if bonus <= 0 || price <= 0 {
		return 0
	}

	budget := float64(bonus)
	levels := 0
	totalCost := 0.0

	for {
		nextLevel := levels + 1
		cansNeeded := nextLevel * nextLevel
		costForLevel := float64(cansNeeded) * price

		if totalCost+costForLevel > budget {
			break
		}

		totalCost += costForLevel
		levels = nextLevel
	}

	return levels
}