/*
## Introduction

Snakes and Ladders is an ancient Indian board game regarded today as a worldwide classic.
It is played by two or more players on a game board with numbered, gridded squares.
A number of "ladders" and "snakes" are pictured on the board, each connecting two specific squares.

## Task

Your task is to create a simple class called `SnakesLadders`. The test cases will call the method
`play(die1, die2)` independently of the state of the game or the player turn. The arguments
`die1` and `die2` are the dice thrown in a turn and are both integers between 1 and 6.
The player will make a number of steps equal to the sum of `die1` and `die2`,
moving one square per step.

## Rules

1. There are two players, and both start off the board on square 0.
2. Player 1 starts and alternates with player 2.
3. You follow the numbers up the board in order from 1 to 100.
4. If the values of both dice are the same, that player will have another turn after the current turn ends.
5. Climb up ladders.
6. Slide down snakes.
7. Land exactly on the last square to win. If you roll too high, your player "bounces" off square 100.
8. If the player rolls a double and lands on the finish square (100), the player wins
   and does not take another turn.

## Returns

Return `"Player n Wins!"` where `n` is the winning player.
Return `"Game over!"` if a move is attempted after any player has won.
Otherwise, return `"Player n is on square x"`.
*/

package kata

type SnakesLadders struct {
	positions     [2]int
	currentPlayer int
	gameOver      bool
}

func NewSnakesLadders() *SnakesLadders {
	return &SnakesLadders{currentPlayer: 0}
}

func (s *SnakesLadders) Play(die1, die2 int) string {
	// TODO: implement
	return ""
}