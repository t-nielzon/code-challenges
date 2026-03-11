package kata

import (
	"fmt"
	"strconv"
	"strings"
)

func StockList(listOfArt []string, listOfCat []string) string {
	if len(listOfArt) == 0 || len(listOfCat) == 0 {
		return ""
	}

	totals := make(map[byte]int)
	for _, art := range listOfArt {
		parts := strings.SplitN(art, " ", 2)
		qty, _ := strconv.Atoi(parts[1])
		totals[parts[0][0]] += qty
	}

	result := make([]string, len(listOfCat))
	for i, cat := range listOfCat {
		result[i] = fmt.Sprintf("(%s : %d)", cat, totals[cat[0]])
	}
	return strings.Join(result, " - ")
}