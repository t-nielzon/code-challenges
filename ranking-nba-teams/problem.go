// Ranking NBA teams
//
// You are given a string with results of NBA teams separated by commas e.g:
//
//   r = `Los Angeles Clippers 104 Dallas Mavericks 88,New York Knicks 101 Atlanta Hawks 112,...`
//
// A team name is composed of one, two or more words built with letters or digits:
// `Atlanta Hawks`, `Philadelphia 76ers`...
//
// Given a string of results and the name of a team (parameter: toFind) your function
// nbaCup will return as a string:
//   - the name of the team followed by `:`
//   - the number of matches won by the team
//   - the number of draws
//   - the number of matches lost by the team
//   - the total number of points scored by the team
//   - the total number of points conceded by the team
//   - and a marks total in the ranking system
//
// Marks: a win marks 3, a draw marks 1, a loss marks 0.
//
// Return format:
//   "Team Name:W=nb of wins;D=nb of draws;L=nb of losses;Scored=nb;Conceded=nb;Points=nb"
//
// Remarks:
//   - The team name "" returns "".
//   - If a team name can't be found you return "Team Name:This team didn't play!".
//   - The scores must be integers. If a score is a float number (abc.xyz...) you return
//     "Error(float number):the concerned string".
//
// Examples:
//   nbaCup(r, "Los Angeles Clippers") -> "Los Angeles Clippers:W=1;D=0;L=1;Scored=204;Conceded=208;Points=3"
//   nbaCup(r, "Boston Celtics")        -> "Boston Celtics:W=1;D=0;L=0;Scored=120;Conceded=100;Points=3"
//   nbaCup(r, "")                      -> ""
//   nbaCup(r, "Boston Celt")           -> "Boston Celt:This team didn't play!"
//   r0 = "New York Knicks 101.12 Atlanta Hawks 112"
//   nbaCup(r0, "Atlanta Hawks")        -> "Error(float number):New York Knicks 101.12 Atlanta Hawks 112"

package kata

func nbaCup(resultSheet string, toFind string) string {
	return ""
}