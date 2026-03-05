package kata

func KingIsInCheck(board [8][8]string) bool {
	var kr, kc int
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			if board[r][c] == "♔" {
				kr, kc = r, c
			}
		}
	}

	// check rook/queen along straight lines
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, d := range dirs {
		for i := 1; ; i++ {
			r, c := kr+d[0]*i, kc+d[1]*i
			if r < 0 || r > 7 || c < 0 || c > 7 {
				break
			}
			p := board[r][c]
			if p == "♜" || p == "♛" {
				return true
			}
			if p != " " {
				break
			}
		}
	}

	// check bishop/queen along diagonals
	diags := [][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	for _, d := range diags {
		for i := 1; ; i++ {
			r, c := kr+d[0]*i, kc+d[1]*i
			if r < 0 || r > 7 || c < 0 || c > 7 {
				break
			}
			p := board[r][c]
			if p == "♝" || p == "♛" {
				return true
			}
			if p != " " {
				break
			}
		}
	}

	// check knight
	knightMoves := [][2]int{{-2, -1}, {-2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {2, -1}, {2, 1}}
	for _, m := range knightMoves {
		r, c := kr+m[0], kc+m[1]
		if r >= 0 && r <= 7 && c >= 0 && c <= 7 && board[r][c] == "♞" {
			return true
		}
	}

	// check pawns — board is from black's perspective, so white pawns attack
	// upward (decreasing row index) meaning they threaten the king from below
	// (row+1) on diagonals
	pawnMoves := [][2]int{{1, -1}, {1, 1}}
	for _, m := range pawnMoves {
		r, c := kr+m[0], kc+m[1]
		if r >= 0 && r <= 7 && c >= 0 && c <= 7 && board[r][c] == "♟" {
			return true
		}
	}

	return false
}