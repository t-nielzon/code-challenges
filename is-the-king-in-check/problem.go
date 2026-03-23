/*
You have to write a function that takes for input a 8x8 chessboard in the form
of a bi-dimensional array of strings of length 1 and returns a boolean
indicating whether the king is in check.

The array will include 64 squares which can contain the following characters:
  - '♔' for the black King
  - '♛' for a white Queen
  - '♝' for a white Bishop
  - '♞' for a white Knight
  - '♜' for a white Rook
  - '♟' for a white Pawn
  - ' ' (a space) if there is no piece on that square

There will always be exactly one king, which is the black king, whereas all
the other pieces are white.
The board is oriented from Black's perspective.
Remember that pawns can only move and take forward.
Also be careful with the pieces' lines of sight.
*/
package kata

func KingIsInCheck(board [8][8]string) bool {
}