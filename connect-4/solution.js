class Connect4 {
  constructor() {
    this.rows = 6;
    this.cols = 7;
    this.board = Array.from({ length: this.rows }, () => Array(this.cols).fill(0));
    this.current = 1;
    this.won = false;
  }

  play(column) {
    if (this.won) return "Game has finished!";

    let row = -1;
    for (let r = this.rows - 1; r >= 0; r--) {
      if (this.board[r][column] === 0) {
        row = r;
        break;
      }
    }
    if (row === -1) return "Column full!";

    this.board[row][column] = this.current;

    if (this.checkWin(row, column, this.current)) {
      const msg = `Player ${this.current} wins!`;
      this.won = true;
      return msg;
    }

    const msg = `Player ${this.current} has a turn`;
    this.current = this.current === 1 ? 2 : 1;
    return msg;
  }

  checkWin(r, c, p) {
    const dirs = [[0, 1], [1, 0], [1, 1], [1, -1]];
    for (const [dr, dc] of dirs) {
      let count = 1;
      for (const sign of [1, -1]) {
        let nr = r + dr * sign, nc = c + dc * sign;
        while (nr >= 0 && nr < this.rows && nc >= 0 && nc < this.cols && this.board[nr][nc] === p) {
          count++;
          nr += dr * sign;
          nc += dc * sign;
        }
      }
      if (count >= 4) return true;
    }
    return false;
  }
}

module.exports = Connect4;