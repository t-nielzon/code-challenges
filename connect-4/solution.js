class Connect4 {
  constructor() {
    this.rows = 6;
    this.cols = 7;
    this.grid = Array.from({ length: this.rows }, () => new Array(this.cols).fill(0));
    this.currentPlayer = 1;
    this.finished = false;
  }

  play(column) {
    if (this.finished) return 'Game has finished!';

    let row = -1;
    for (let r = this.rows - 1; r >= 0; r--) {
      if (this.grid[r][column] === 0) {
        row = r;
        break;
      }
    }

    if (row === -1) return 'Column full!';

    const player = this.currentPlayer;
    this.grid[row][column] = player;

    if (this.checkWin(row, column, player)) {
      this.finished = true;
      return `Player ${player} wins!`;
    }

    this.currentPlayer = player === 1 ? 2 : 1;
    return `Player ${player} has a turn`;
  }

  checkWin(row, col, player) {
    const directions = [
      [0, 1],
      [1, 0],
      [1, 1],
      [1, -1],
    ];

    for (const [dr, dc] of directions) {
      let count = 1;

      for (let dir = -1; dir <= 1; dir += 2) {
        let r = row + dr * dir;
        let c = col + dc * dir;
        while (
          r >= 0 && r < this.rows &&
          c >= 0 && c < this.cols &&
          this.grid[r][c] === player
        ) {
          count++;
          r += dr * dir;
          c += dc * dir;
        }
      }

      if (count >= 4) return true;
    }

    return false;
  }
}

module.exports = Connect4;