package kata

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func NbaCup(resultSheet string, toFind string) string {
	if toFind == "" {
		return ""
	}

	floatRe := regexp.MustCompile(`\d+\.\d+`)
	intRe := regexp.MustCompile(`^\d+$`)

	matches := strings.Split(resultSheet, ",")
	wins, draws, losses := 0, 0, 0
	scored, conceded := 0, 0
	found := false

	for _, m := range matches {
		m = strings.TrimSpace(m)
		if m == "" {
			continue
		}

		if floatRe.MatchString(m) {
			return "Error(float number):" + m
		}

		tokens := strings.Fields(m)
		intPositions := []int{}
		for i, t := range tokens {
			if intRe.MatchString(t) {
				intPositions = append(intPositions, i)
			}
		}

		if len(intPositions) != 2 {
			continue
		}

		p1, p2 := intPositions[0], intPositions[1]
		teamA := strings.Join(tokens[:p1], " ")
		scoreA, _ := strconv.Atoi(tokens[p1])
		teamB := strings.Join(tokens[p1+1:p2], " ")
		scoreB, _ := strconv.Atoi(tokens[p2])

		if teamA == toFind {
			found = true
			scored += scoreA
			conceded += scoreB
			switch {
			case scoreA > scoreB:
				wins++
			case scoreA == scoreB:
				draws++
			default:
				losses++
			}
		}
		if teamB == toFind {
			found = true
			scored += scoreB
			conceded += scoreA
			switch {
			case scoreB > scoreA:
				wins++
			case scoreA == scoreB:
				draws++
			default:
				losses++
			}
		}
	}

	if !found {
		return toFind + ":This team didn't play!"
	}

	points := wins*3 + draws
	return fmt.Sprintf("%s:W=%d;D=%d;L=%d;Scored=%d;Conceded=%d;Points=%d",
		toFind, wins, draws, losses, scored, conceded, points)
}