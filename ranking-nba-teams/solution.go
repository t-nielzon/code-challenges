package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func nbaCup(results string, teamName string) string {
	if teamName == "" {
		return ""
	}

	games := strings.Split(results, ",")

	type TeamStats struct {
		wins     int
		draws    int
		losses   int
		scored   int
		conceded int
	}

	stats := &TeamStats{}
	found := false

	for _, game := range games {
		game = strings.TrimSpace(game)
		if game == "" {
			continue
		}

		team1, score1, team2, score2, errMsg := parseGame(game)
		if errMsg != "" {
			return errMsg
		}

		if team1 == teamName {
			found = true
			if score1 > score2 {
				stats.wins++
			} else if score1 < score2 {
				stats.losses++
			} else {
				stats.draws++
			}
			stats.scored += score1
			stats.conceded += score2
		} else if team2 == teamName {
			found = true
			if score2 > score1 {
				stats.wins++
			} else if score2 < score1 {
				stats.losses++
			} else {
				stats.draws++
			}
			stats.scored += score2
			stats.conceded += score1
		}
	}

	if !found {
		return fmt.Sprintf("%s:This team didn't play!", teamName)
	}

	points := stats.wins*3 + stats.draws*1

	return fmt.Sprintf("%s:W=%d;D=%d;L=%d;Scored=%d;Conceded=%d;Points=%d",
		teamName, stats.wins, stats.draws, stats.losses, stats.scored, stats.conceded, points)
}

func parseGame(game string) (string, int, string, int, string) {
	parts := strings.Fields(game)

	var numberIndices []int
	var scores []int
	numberPattern := regexp.MustCompile(`^(\d+)(\.\d+)?$`)

	for i, part := range parts {
		if match := numberPattern.FindStringSubmatch(part); match != nil {
			numberIndices = append(numberIndices, i)

			if match[2] != "" {
				return "", 0, "", 0, fmt.Sprintf("Error(float number):%s", game)
			}

			score, _ := strconv.Atoi(match[1])
			scores = append(scores, score)
		}
	}

	if len(scores) != 2 {
		return "", 0, "", 0, ""
	}

	firstScoreIdx := numberIndices[0]
	secondScoreIdx := numberIndices[1]

	team1 := strings.Join(parts[:firstScoreIdx], " ")
	team2 := strings.Join(parts[firstScoreIdx+1:secondScoreIdx], " ")

	return team1, scores[0], team2, scores[1], ""
}