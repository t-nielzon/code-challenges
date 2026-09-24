package main

import "fmt"

type SnakesLadders struct {
	player1Pos    int
	player2Pos    int
	currentPlayer int
	gameWon       bool
	snakes        map[int]int
	ladders       map[int]int
}

func NewSnakesLadders() *SnakesLadders {
	return &SnakesLadders{
		player1Pos:    0,
		player2Pos:    0,
		currentPlayer: 1,
		gameWon:       false,
		snakes: map[int]int{
			17: 7, 54: 34, 62: 19, 88: 24, 93: 73, 95: 75, 98: 79, 99: 40,
		},
		ladders: map[int]int{
			2: 38, 7: 14, 15: 26, 21: 42, 28: 84, 51: 67, 72: 91, 78: 98, 87: 94,
		},
	}
}

func (sl *SnakesLadders) Play(die1, die2 int) string {
	if sl.gameWon {
		return "Game over!"
	}

	steps := die1 + die2
	var pos int

	if sl.currentPlayer == 1 {
		pos = sl.player1Pos
	} else {
		pos = sl.player2Pos
	}

	movedPlayer := sl.currentPlayer

	pos += steps

	// bounce off 100
	if pos > 100 {
		pos = 100 - (pos - 100)
	}

	// apply snake or ladder
	if snake, ok := sl.snakes[pos]; ok {
		pos = snake
	} else if ladder, ok := sl.ladders[pos]; ok {
		pos = ladder
	}

	if sl.currentPlayer == 1 {
		sl.player1Pos = pos
	} else {
		sl.player2Pos = pos
	}

	if pos == 100 {
		sl.gameWon = true
		return fmt.Sprintf("Player %d Wins!", movedPlayer)
	}

	if die1 != die2 {
		sl.currentPlayer = 3 - sl.currentPlayer
	}

	return fmt.Sprintf("Player %d is on square %d", movedPlayer, pos)
}