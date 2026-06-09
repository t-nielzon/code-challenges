// ## Introduction
//
// Snakes and Ladders is an ancient Indian board game regarded today as a
// worldwide classic. It is played by two or more players on a game board with
// numbered, gridded squares. A number of "ladders" and "snakes" are pictured on
// the board, each connecting two specific squares.
//
// ## Task
//
// Your task is to create a simple class called `SnakesLadders`. The test cases
// will call the method `play(die1, die2)` independently of the state of the game
// or the player turn. The arguments `die1` and `die2` are the dice thrown in a
// turn and are both integers between 1 and 6. The player will make a number of
// steps equal to the sum of `die1` and `die2`, moving one square per step.
//
// ## Rules
//
// 1. There are two players, and both start off the board on square 0.
// 2. Player 1 starts and alternates with player 2.
// 3. You follow the numbers up the board in order from 1 to 100.
// 4. If the values of both dice are the same, that player will have another turn
//    after the current turn ends.
// 5. Climb up ladders. If you land exactly on a square that shows the bottom of a
//    ladder, you move all the way up to the top (even if you roll a double).
// 6. Slide down snakes. If you land exactly on the top of a snake, you slide all
//    the way down to the bottom (even if you roll a double).
// 7. Land exactly on the last square to win. If you roll too high, your player
//    "bounces" off square 100 and continues moving backward for the remaining
//    steps. You can only win by rolling the exact number needed.
// 8. If the player rolls a double and lands on the finish square (100) after
//    taking all steps, the player wins and does not take another turn.
//
// ## Returns
//
// Return "Player n Wins!" where n is the winning player who has landed on
// square 100 after taking all steps in their turn.
// Return "Game over!" if a move is attempted after any player has won.
// Otherwise, return "Player n is on square x", where n is the current player and
// x is the square they are currently on.

package kata

type SnakesLadders struct {
	positions [2]int
	turn      int
	gameOver  bool
}

func NewSnakesLadders() *SnakesLadders {
	return &SnakesLadders{}
}

func (s *SnakesLadders) play(die1, die2 int) string {
	return ""
}