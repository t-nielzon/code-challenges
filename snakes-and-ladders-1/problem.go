package main

/*
Snakes and Ladders is an ancient Indian board game regarded today as a worldwide classic.
It is played by two or more players on a game board with numbered, gridded squares.
A number of "ladders" and "snakes" are pictured on the board, each connecting two specific squares.

Task: Create a simple class called SnakesLadders. The test cases will call the method Play(die1, die2)
independently of the state of the game or the player turn. The arguments die1 and die2 are the dice
thrown in a turn and are both integers between 1 and 6. The player will make a number of steps equal
to the sum of die1 and die2, moving one square per step.

Rules:
1. Two players start at square 0.
2. Player 1 starts and alternates with player 2.
3. Follow numbers 1 to 100.
4. If both dice values are the same, the player gets another turn.
5. Climb up ladders: if landing exactly on a ladder bottom, move to the ladder top.
6. Slide down snakes: if landing exactly on a snake top, move to the snake bottom.
7. Land exactly on 100 to win. If rolling too high, bounce off 100 and move backward.
8. If rolling doubles and landing on 100, the player wins without another turn.

Returns:
- "Player n Wins!" when a player reaches square 100.
- "Game over!" if a move is attempted after a player has won.
- "Player n is on square x" otherwise, showing current player and their square.
*/

type SnakesLadders struct {
	// add your fields here
}

func NewSnakesLadders() *SnakesLadders {
	// implement me
	return nil
}

func (sl *SnakesLadders) Play(die1, die2 int) string {
	// implement me
	return ""
}