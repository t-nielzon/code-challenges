package kata

import "math"

func NbMonths(startPriceOld, startPriceNew, savingPerMonth int, percentLossByMonth float64) [2]int {
	oldPrice := float64(startPriceOld)
	newPrice := float64(startPriceNew)
	saving := 0.0
	months := 0
	loss := percentLossByMonth

	for oldPrice+saving < newPrice {
		months++
		if months%2 == 0 {
			loss += 0.5
		}
		oldPrice -= oldPrice * loss / 100
		newPrice -= newPrice * loss / 100
		saving += float64(savingPerMonth)
	}

	leftover := oldPrice + saving - newPrice
	return [2]int{months, int(math.Round(leftover))}
}