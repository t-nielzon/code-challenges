package kata

import "strings"

func PingPong(sounds string) string {
	parts := strings.Split(sounds, "-")

	scores := map[string]int{"ping": 0, "pong": 0}
	var server string
	var lastHitter string
	var lastBadHitter string
	inRally := false

	for _, s := range parts {
		if s == "ping" || s == "pong" {
			if !inRally {
				server = s
				inRally = true
			}
			lastHitter = s
		} else {
			if inRally {
				inRally = false
				lastBadHitter = lastHitter
				winner := "ping"
				if lastHitter == "ping" {
					winner = "pong"
				}
				if server == winner {
					scores[winner]++
				}
			}
		}
	}

	if scores["ping"] > scores["pong"] {
		return "ping"
	}
	if scores["pong"] > scores["ping"] {
		return "pong"
	}

	if lastBadHitter == "ping" {
		return "pong"
	}
	return "ping"
}