package main

func Beeramid(bonus float32, price float32) int {
	if bonus <= 0 || price <= 0 {
		return 0
	}

	totalCost := float32(0)

	for level := 1; ; level++ {
		cansInLevel := level * level
		costOfLevel := float32(cansInLevel) * price

		if totalCost+costOfLevel > bonus {
			return level - 1
		}

		totalCost += costOfLevel
	}
}