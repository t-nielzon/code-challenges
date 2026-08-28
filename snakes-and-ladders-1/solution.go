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
			17: 7,
			54: 34,
			62: 18,
			87: 24,
			93: 73,
			95: 75,
			98: 79,
		},
		ladders: map[int]int{
			1:  38,
			4:  14,
			9:  31,
			21: 42,
			28: 84,
			51: 67,
			72: 91,
			80: 99,
		},
	}
}

func (sl *SnakesLadders) Play(die1, die2 int) string {
	if sl.gameWon {
		return "Game over!"
	}

	player := sl.currentPlayer
	var pos int

	if player == 1 {
		pos = sl.player1Pos
	} else {
		pos = sl.player2Pos
	}

	pos += die1 + die2

	if pos > 100 {
		pos = 200 - pos
	}

	if ladder, exists := sl.ladders[pos]; exists {
		pos = ladder
	}

	if snake, exists := sl.snakes[pos]; exists {
		pos = snake
	}

	if player == 1 {
		sl.player1Pos = pos
	} else {
		sl.player2Pos = pos
	}

	if pos == 100 {
		sl.gameWon = true
		return fmt.Sprintf("Player %d Wins!", player)
	}

	if die1 != die2 {
		sl.currentPlayer = 3 - player
	}

	return fmt.Sprintf("Player %d is on square %d", player, pos)
}