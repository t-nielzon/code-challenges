package kata

func KingIsInCheck(board [8][8]string) bool {
	// find the king
	var kr, kc int
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			if board[r][c] == "♔" {
				kr, kc = r, c
			}
		}
	}

	// check knight attacks
	knightMoves := [][2]int{
		{-2, -1}, {-2, 1}, {-1, -2}, {-1, 2},
		{1, -2}, {1, 2}, {2, -1}, {2, 1},
	}
	for _, m := range knightMoves {
		r, c := kr+m[0], kc+m[1]
		if r >= 0 && r < 8 && c >= 0 && c < 8 && board[r][c] == "♞" {
			return true
		}
	}

	// check pawn attacks
	// board is from black's perspective, so pawns attack "downward" (increasing row)
	// white pawns move up from white's perspective, which is down from black's perspective
	// pawns capture diagonally forward (toward the king's side = row+1 from black's top)
	// actually: black is at top (row 0). white pawns move upward toward row 0.
	// so white pawns at row kr+1 can capture diagonally to kr
	for _, dc := range []int{-1, 1} {
		r, c := kr+1, kc+dc
		if r >= 0 && r < 8 && c >= 0 && c < 8 && board[r][c] == "♟" {
			return true
		}
	}

	// check straight lines for rook/queen
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, d := range dirs {
		if castRay(board, kr, kc, d[0], d[1], "♜", "♛") {
			return true
		}
	}

	// check diagonals for bishop/queen
	diags := [][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	for _, d := range diags {
		if castRay(board, kr, kc, d[0], d[1], "♝", "♛") {
			return true
		}
	}

	return false
}

func castRay(board [8][8]string, kr, kc, dr, dc int, piece1, piece2 string) bool {
	r, c := kr+dr, kc+dc
	for r >= 0 && r < 8 && c >= 0 && c < 8 {
		if board[r][c] != " " {
			return board[r][c] == piece1 || board[r][c] == piece2
		}
		r += dr
		c += dc
	}
	return false
}