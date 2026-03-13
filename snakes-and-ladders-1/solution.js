class SnakesLadders {
  constructor() {
    this.positions = [0, 0];
    this.currentPlayer = 0;
    this.winner = false;

    this.board = {
      2: 38, 7: 14, 8: 31, 15: 26, 21: 42, 28: 84, 36: 44,
      51: 67, 71: 91, 78: 98, 87: 94,
      16: 6, 46: 25, 49: 11, 62: 19, 64: 60, 74: 53,
      89: 68, 92: 88, 95: 75, 99: 80
    };
  }

  play(die1, die2) {
    if (this.winner) return 'Game over!';

    const player = this.currentPlayer;
    let pos = this.positions[player] + die1 + die2;

    if (pos > 100) {
      pos = 200 - pos;
    }

    if (this.board[pos] !== undefined) {
      pos = this.board[pos];
    }

    this.positions[player] = pos;

    if (pos === 100) {
      this.winner = true;
      return `Player ${player + 1} Wins!`;
    }

    if (die1 !== die2) {
      this.currentPlayer = 1 - this.currentPlayer;
    }

    return `Player ${player + 1} is on square ${pos}`;
  }
}