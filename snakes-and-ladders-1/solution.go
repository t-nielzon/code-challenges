package kata

import "fmt"

type SnakesLadders struct {
	players       [2]int
	currentPlayer int
	gameOver      bool
}

var ladders = map[int]int{
	2:  38,
	7:  14,
	8:  31,
	15: 26,
	21: 42,
	28: 84,
	36: 44,
	51: 67,
	71: 91,
	78: 98,
	87: 94,
}

var snakes = map[int]int{
	16:  6,
	46:  25,
	49:  11,
	62:  19,
	64:  60,
	74:  53,
	89:  68,
	92:  88,
	95:  75,
	99:  80,
}

func NewSnakesLadders() *SnakesLadders {
	return &SnakesLadders{
		players:       [2]int{0, 0},
		currentPlayer: 0,
		gameOver:      false,
	}
}

func (s *SnakesLadders) Play(die1, die2 int) string {
	if s.gameOver {
		return "Game over!"
	}

	move := die1 + die2
	pos := s.players[s.currentPlayer] + move

	if pos > 100 {
		pos = 100 - (pos - 100)
	}

	if dest, ok := ladders[pos]; ok {
		pos = dest
	} else if dest, ok := snakes[pos]; ok {
		pos = dest
	}

	s.players[s.currentPlayer] = pos

	playerNum := s.currentPlayer + 1

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