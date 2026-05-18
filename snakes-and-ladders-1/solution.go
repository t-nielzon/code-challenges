package kata

import "fmt"

type SnakesLadders struct {
	positions     [2]int
	currentPlayer int
	gameOver      bool
}

func NewSnakesLadders() *SnakesLadders {
	return &SnakesLadders{currentPlayer: 0}
}

var ladders = map[int]int{
	2: 38, 7: 14, 8: 31, 15: 26, 21: 42, 28: 84,
	36: 44, 51: 67, 71: 91, 78: 98, 87: 94,
}

var snakes = map[int]int{
	16: 6, 46: 25, 49: 11, 62: 19, 64: 60,
	74: 53, 89: 68, 92: 88, 95: 75, 99: 80,
}

func (s *SnakesLadders) Play(die1, die2 int) string {
	if s.gameOver {
		return "Game over!"
	}

	p := s.currentPlayer
	pos := s.positions[p] + die1 + die2
	if pos > 100 {
		pos = 200 - pos
	}
	if v, ok := ladders[pos]; ok {
		pos = v
	} else if v, ok := snakes[pos]; ok {
		pos = v
	}
	s.positions[p] = pos

	playerNum := p + 1

	if pos == 100 {
		s.gameOver = true
		return fmt.Sprintf("Player %d Wins!", playerNum)
	}

	result := fmt.Sprintf("Player %d is on square %d", playerNum, pos)

	if die1 != die2 {
		s.currentPlayer = 1 - s.currentPlayer
	}

	return result
}