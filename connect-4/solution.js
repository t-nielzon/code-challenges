class Connect4 {
  constructor() {
    this.board = Array(6).fill(null).map(() => Array(7).fill(0));
    this.currentPlayer = 1;
    this.gameWon = false;
  }

  play(column) {
    if (this.gameWon) {
      return "Game has finished!";
    }

    let row = -1;
    for (let i = 5; i >= 0; i--) {
      if (this.board[i][column] === 0) {
        row = i;
        break;
      }
    }

    if (row === -1) {
      return "Column full!";
    }

    this.board[row][column] = this.currentPlayer;

    if (this.checkWin(row, column)) {
      this.gameWon = true;
      return `Player ${this.currentPlayer} wins!`;
    }

    const message = `Player ${this.currentPlayer} has a turn`;
    this.currentPlayer = this.currentPlayer === 1 ? 2 : 1;
    return message;
  }

  checkWin(row, col) {
    const player = this.board[row][col];

    // horizontal
    let count = 1;
    for (let c = col - 1; c >= 0 && this.board[row][c] === player; c--) count++;
    for (let c = col + 1; c < 7 && this.board[row][c] === player; c++) count++;
    if (count >= 4) return true;

    // vertical
    count = 1;
    for (let r = row - 1; r >= 0 && this.board[r][col] === player; r--) count++;
    for (let r = row + 1; r < 6 && this.board[r][col] === player; r++) count++;
    if (count >= 4) return true;

    // diagonal (top-left to bottom-right)
    count = 1;
    for (let r = row - 1, c = col - 1; r >= 0 && c >= 0 && this.board[r][c] === player; r--, c--) count++;
    for (let r = row + 1, c = col + 1; r < 6 && c < 7 && this.board[r][c] === player; r++, c++) count++;
    if (count >= 4) return true;

    // diagonal (top-right to bottom-left)
    count = 1;
    for (let r = row - 1, c = col + 1; r >= 0 && c < 7 && this.board[r][c] === player; r--, c++) count++;
    for (let r = row + 1, c = col - 1; r < 6 && c >= 0 && this.board[r][c] === player; r++, c--) count++;
    if (count >= 4) return true;

    return false;
  }
}