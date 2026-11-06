package kata

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func NbaCup(results string, toFind string) string {
	if toFind == "" {
		return ""
	}

	scorePattern := regexp.MustCompile(`\d+(?:\.\d+)?`)

	games := strings.Split(results, ",")

	wins := 0
	draws := 0
	losses := 0
	scored := 0
	conceded := 0
	found := false

	for _, game := range games {
		game = strings.TrimSpace(game)
		if game == "" {
			continue
		}

		scoreMatches := scorePattern.FindAllStringIndex(game, -1)
		if len(scoreMatches) != 2 {
			continue
		}

		score1Str := game[scoreMatches[0][0]:scoreMatches[0][1]]
		score2Str := game[scoreMatches[1][0]:scoreMatches[1][1]]

		if strings.Contains(score1Str, ".") || strings.Contains(score2Str, ".") {
			return fmt.Sprintf("Error(float number):%s", game)
		}

		score1, _ := strconv.Atoi(score1Str)
		score2, _ := strconv.Atoi(score2Str)

		team1 := strings.TrimSpace(game[:scoreMatches[0][0]])
		team2 := strings.TrimSpace(game[scoreMatches[0][1]:scoreMatches[1][0]])

		if team1 == toFind {
			found = true
			scored += score1
			conceded += score2

			if score1 > score2 {
				wins++
			} else if score1 < score2 {
				losses++
			} else {
				draws++
			}
		} else if team2 == toFind {
			found = true
			scored += score2
			conceded += score1

			if score2 > score1 {
				wins++
			} else if score2 < score1 {
				losses++
			} else {
				draws++
			}
		}
	}

	if !found {
		return fmt.Sprintf("%s:This team didn't play!", toFind)
	}

	points := wins*3 + draws*1

	return fmt.Sprintf("%s:W=%d;D=%d;L=%d;Scored=%d;Conceded=%d;Points=%d",
		toFind, wins, draws, losses, scored, conceded, points)
}