package kata

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var matchRe = regexp.MustCompile(`^(.+?)\s+(\d+(?:\.\d+)?)\s+(.+?)\s+(\d+(?:\.\d+)?)$`)

func nbaCup(resultSheet string, toFind string) string {
	if toFind == "" {
		return ""
	}

	found := false
	wins, draws, losses, scored, conceded := 0, 0, 0, 0, 0

	for _, raw := range strings.Split(resultSheet, ",") {
		chunk := strings.TrimSpace(raw)
		if chunk == "" {
			continue
		}

		m := matchRe.FindStringSubmatch(chunk)
		if m == nil {
			continue
		}

		homeName, homeScoreStr := m[1], m[2]
		awayName, awayScoreStr := m[3], m[4]

		// a float score invalidates the whole reading for the concerned string
		if strings.Contains(homeScoreStr, ".") || strings.Contains(awayScoreStr, ".") {
			return "Error(float number):" + chunk
		}

		homeScore, _ := strconv.Atoi(homeScoreStr)
		awayScore, _ := strconv.Atoi(awayScoreStr)

		var teamScore, oppScore int
		switch toFind {
		case homeName:
			teamScore, oppScore = homeScore, awayScore
		case awayName:
			teamScore, oppScore = awayScore, homeScore
		default:
			continue
		}

		found = true
		scored += teamScore
		conceded += oppScore
		switch {
		case teamScore > oppScore:
			wins++
		case teamScore == oppScore:
			draws++
		default:
			losses++
		}
	}

	if !found {
		return toFind + ":This team didn't play!"
	}

	points := wins*3 + draws
	return fmt.Sprintf("%s:W=%d;D=%d;L=%d;Scored=%d;Conceded=%d;Points=%d",
		toFind, wins, draws, losses, scored, conceded, points)
}