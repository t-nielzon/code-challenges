package main

import "fmt"

type SnakesLadders struct {
	player1Pos int
	player2Pos int
	current    int
	finished   bool
}

func New() *SnakesLadders {
	return &SnakesLadders{
		player1Pos: 0,
		player2Pos: 0,
		current:    1,
		finished:   false,
	}
}

var ladders = map[int]int{
	1:  38,
	4:  14,
	9:  31,
	21: 42,
	28: 84,
	51: 67,
	72: 91,
	80: 99,
}

var snakes = map[int]int{
	17: 7,
	54: 34,
	62: 18,
	88: 24,
	95: 75,
	98: 79,
}

func (sl *SnakesLadders) Play(die1, die2 int) string {
	if sl.finished {
		return "Game over!"
	}

	var pos int
	if sl.current == 1 {
		pos = sl.player1Pos
	} else {
		pos = sl.player2Pos
	}

	steps := die1 + die2
	pos += steps

	if pos > 100 {
		pos = 100 - (pos - 100)
	}

	if newPos, exists := ladders[pos]; exists {
		pos = newPos
	} else if newPos, exists := snakes[pos]; exists {
		pos = newPos
	}

	if sl.current == 1 {
		sl.player1Pos = pos
	} else {
		sl.player2Pos = pos
	}

	if pos == 100 {
		sl.finished = true
		return fmt.Sprintf("Player %d Wins!", sl.current)
	}

	result := fmt.Sprintf("Player %d is on square %d", sl.current, pos)

	if die1 != die2 {
		if sl.current == 1 {
			sl.current = 2
		} else {
			sl.current = 1
		}
	}

	return result
}