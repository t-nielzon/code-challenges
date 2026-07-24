package main

import "strings"

func PingPong(sounds string) string {
	soundList := strings.Split(sounds, "-")
	
	pingScore := 0
	pongScore := 0
	
	var currentServer string
	var lastValidHitter string
	var finalBadHitter string
	
	for _, sound := range soundList {
		if sound == "ping" || sound == "pong" {
			if currentServer == "" {
				currentServer = sound
			}
			lastValidHitter = sound
		} else {
			// bad noise marks end of rally
			if currentServer != "" {
				if currentServer != lastValidHitter {
					// server wins (opponent made the bad shot)
					if currentServer == "ping" {
						pingScore++
					} else {
						pongScore++
					}
				}
			}
			// track final bad hitter for tie-breaking
			finalBadHitter = lastValidHitter
			// reset for next rally
			currentServer = ""
			lastValidHitter = ""
		}
	}
	
	// determine winner
	if pingScore > pongScore {
		return "ping"
	} else if pongScore > pingScore {
		return "pong"
	} else {
		// tie: winner is who didn't hit the final bad shot
		if finalBadHitter == "ping" {
			return "pong"
		}
		return "ping"
	}
}