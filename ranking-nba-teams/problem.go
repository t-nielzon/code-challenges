/*
You are given a string with results of NBA teams separated by commas e.g:

r =
`Los Angeles Clippers 104 Dallas Mavericks 88,New York Knicks 101 Atlanta Hawks 112,
 Indiana Pacers 103 Memphis Grizzlies 112,Los Angeles Clippers 100 Boston Celtics 120`.

A team name is composed of one, two or more words built with letters or digits:
`Atlanta Hawks`, `Philadelphia 76ers`...

Given a string of results and the name of a team (parameter : to_find) your function
nbaCup will return as a string:
- the name of the team followed by `:`
- the number of matches won
- the number of draws
- the number of matches lost
- the total number of points scored
- the total number of points conceded
and a marks total:
- a team marks 3 if it is a win
- a team marks 1 if it is a draw
- a team marks 0 if it is a loss

Return format:
"Team Name:W=nb of wins;D=nb of draws;L=nb of losses;Scored=nb;Conceded=nb;Points=nb"

Remarks:
- The team name "" returns "".
- If a team name can't be found, return "Team Name:This team didn't play!"
- If a score is a float number, return "Error(float number):the concerned string"
*/

package kata

func NbaCup(resultSheet string, toFind string) string {
	// your code here
	return ""
}