package kata

import "fmt"

type SnakesLadders struct {
	positions [2]int
	turn      int
	gameOver  bool
}

func NewSnakesLadders() *SnakesLadders {
	return &SnakesLadders{}
}

// jumps maps the bottom of a ladder (or top of a snake) to its destination square.
var jumps = map[int]int{
	// ladders
	2: 38, 7: 14, 8: 31, 15: 26, 21: 42, 28: 84,
	36: 44, 51: 67, 71: 91, 78: 98, 87: 94,
	// snakes
	16: 6, 46: 25, 49: 11, 62: 19, 64: 60, 74: 53,
	89: 68, 92: 88, 95: 75, 99: 80,
}

func (s *SnakesLadders) play(die1, die2 int) string {
	if s.gameOver {
		return "Game over!"
	}

	player := s.turn
	pos := s.positions[player] + die1 + die2

	// bounce backward off the final square when overshooting
	if pos > 100 {
		pos = 200 - pos
	}

	// landing exactly on the foot of a ladder or head of a snake relocates the piece
	if dest, ok := jumps[pos]; ok {
		pos = dest
	}

	s.positions[player] = pos

	if pos == 100 {
		s.gameOver = true
		return fmt.Sprintf("Player %d Wins!", player+1)
	}

	// doubles grant the same player another turn; otherwise pass to the other player
	if die1 != die2 {
		s.turn = 1 - s.turn
	}

	return fmt.Sprintf("Player %d is on square %d", player+1, pos)
}