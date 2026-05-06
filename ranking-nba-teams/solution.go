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

	intRe := regexp.MustCompile(`^\d+$`)
	floatRe := regexp.MustCompile(`^\d+\.\d+$`)

	rawMatches := strings.Split(resultSheet, ",")
	cleaned := make([]string, 0, len(rawMatches))
	for _, m := range rawMatches {
		m = strings.TrimSpace(m)
		if m == "" {
			continue
		}
		cleaned = append(cleaned, m)
	}

	for _, m := range cleaned {
		for _, tok := range strings.Fields(m) {
			if floatRe.MatchString(tok) {
				return "Error(float number):" + m
			}
		}
	}

	wins, draws, losses, scored, conceded := 0, 0, 0, 0, 0
	found := false

	for _, m := range cleaned {
		tokens := strings.Fields(m)
		var idx []int
		for i, t := range tokens {
			if intRe.MatchString(t) {
				idx = append(idx, i)
			}
		}
		if len(idx) < 2 {
			continue
		}

		i1, i2 := idx[0], idx[len(idx)-1]
		team1 := strings.Join(tokens[:i1], " ")
		team2 := strings.Join(tokens[i1+1:i2], " ")
		score1, _ := strconv.Atoi(tokens[i1])
		score2, _ := strconv.Atoi(tokens[i2])

		var ts, os int
		switch toFind {
		case team1:
			ts, os = score1, score2
		case team2:
			ts, os = score2, score1
		default:
			continue
		}

		found = true
		scored += ts
		conceded += os
		switch {
		case ts > os:
			wins++
		case ts == os:
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