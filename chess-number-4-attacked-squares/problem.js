// Problem stub: problem.js
/*
Chess #4: Attacked Squares

A key skill in the game of chess is knowing what squares you are attacking,
and what squares your opponent is attacking.

Input: An 8x8 2D array representing a chess board from White's perspective.
Each square can contain:
- " " = Nothing
- "P" or "p" = Pawn (white or black)
- "N" or "n" = Knight
- "B" or "b" = Bishop
- "R" or "r" = Rook
- "Q" or "q" = Queen
- "K" or "k" = King
- Uppercase = White piece, lowercase = Black piece

Output: Return a 2D array where each square shows:
- "  " = Square is not attacked
- "W " = Square is attacked by white
- " B" = Square is attacked by black
- "WB" = Square is attacked by white and black

Rules:
- Pieces do not attack the square they occupy
- Pieces attack squares occupied by another piece of the same color (defending)
- White pawns move towards rank 8 (index 0), black pawns towards rank 1 (index 7)
- Pawns only attack diagonally
- Do not worry about illegal positions or en passant
*/

function attackedSquares(board) {
  
}