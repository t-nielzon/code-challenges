// Snakes and Ladders is an ancient Indian board game regarded today as a worldwide classic.
// It is played by two or more players on a game board with numbered, gridded squares.
// A number of "ladders" and "snakes" are pictured on the board, each connecting two specific squares.
//
// Task: Create a simple class called SnakesLadders.
// The test cases will call the method Play(die1, die2) independently of the game state.
// The arguments die1 and die2 are the dice thrown in a turn (integers between 1 and 6).
// The player moves a number of steps equal to die1 + die2.
//
// Rules:
// 1. Two players, both start at square 0.
// 2. Player 1 starts and alternates with player 2.
// 3. Follow numbers up the board from 1 to 100.
// 4. If both dice are the same, that player gets another turn.
// 5. Climb up ladders to move ahead faster.
// 6. Slide down snakes to move back on the board.
// 7. Land exactly on 100 to win. If you roll too high, bounce back.
// 8. If player rolls double and lands on 100, they win without another turn.
//
// Returns:
// - "Player n Wins!" when player n lands on square 100.
// - "Game over!" if a move is attempted after the game has ended.
// - "Player n is on square x" for the current state.

package main

type SnakesLadders struct {
}

func New() *SnakesLadders {
	return &SnakesLadders{}
}

func (sl *SnakesLadders) Play(die1, die2 int) string {
	// TODO: Implement the Play method
	return ""
}