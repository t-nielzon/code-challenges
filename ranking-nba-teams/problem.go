package main

/*
Ranking NBA teams

You are given a string with results of NBA teams separated by commas.
A team name is composed of one, two or more words built with letters or digits.

Given a string of results and the name of a team, your function nbaCup
will return team statistics including wins, draws, losses, points scored,
points conceded, and ranking points.

Return format:
"Team Name:W=nb of wins;D=nb of draws;L=nb of losses;Scored=nb;Conceded=nb;Points=nb"

Ranking system:
- 3 points for a win
- 1 point for a draw
- 0 points for a loss

Special cases:
- Empty team name returns ""
- Team not found returns "Team Name:This team didn't play!"
- Float scores return "Error(float number):the concerned string"
*/
func nbaCup(results string, teamName string) string {
	return ""
}