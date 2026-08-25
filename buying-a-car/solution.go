package kata

import "math"

func nbMonths(startPriceOld int, startPriceNew int, savingPerMonth int, percentLossPerMonth float64) [2]int {
	oldCarPrice := float64(startPriceOld)
	newCarPrice := float64(startPriceNew)
	savings := 0.0
	month := 0
	percentLoss := percentLossPerMonth
	
	if oldCarPrice >= newCarPrice {
		return [2]int{0, int(math.Round(oldCarPrice - newCarPrice))}
	}
	
	for {
		month++
		
		if month%2 == 0 {
			percentLoss += 0.5
		}
		
		savings += float64(savingPerMonth)
		
		oldCarPrice *= (1 - percentLoss/100)
		newCarPrice *= (1 - percentLoss/100)
		
		available := oldCarPrice + savings - newCarPrice
		if available >= 0 {
			return [2]int{month, int(math.Round(available))}
		}
	}
}