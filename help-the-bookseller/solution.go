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
	parts := make([]string, 0, len(listOfCat))
	for _, cat := range listOfCat {
		total := 0
		for _, art := range listOfArt {
			fields := strings.Fields(art)
			if len(fields) < 2 || len(fields[0]) == 0 {
				continue
			}
			if string(fields[0][0]) == cat {
				n, err := strconv.Atoi(fields[1])
				if err == nil {
					total += n
				}
			}
		}
		parts = append(parts, fmt.Sprintf("(%s : %d)", cat, total))
	}
	return strings.Join(parts, " - ")
}