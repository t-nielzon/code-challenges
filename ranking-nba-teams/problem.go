package kata

// Ranking NBA teams
//
// You are given a string with results of NBA teams separated by commas.
// A team name is composed of one, two or more words built with letters or digits.
//
// Given a string of results and the name of a team, your function NbaCup will return:
// - the name of the team followed by `:` and
// - the number of matches won by the team
// - the number of draws
// - the number of matches lost by the team
// - the total number of points scored by the team
// - the total number of points conceded by the team
// - and a kind of marks in our ranking system
//
// Ranking points:
// - a team marks `3` if it is a win
// - a team marks `1` if it is a draw
// - a team marks `0` if it is a loss
//
// The return format is:
// "Team Name:W=nb of wins;D=nb of draws;L=nb of losses;Scored=nb;Conceded=nb;Points=nb"
//
// Remarks:
// - The team name "" returns "".
// - If a team name can't be found, return "Team Name:This team didn't play!"
// - The scores must be integers. If a score is a float number, return:
//   "Error(float number):the concerned string"
func NbaCup(results string, toFind string) string {
	return ""
}