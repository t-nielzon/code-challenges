class SnakesLadders {
  constructor() {
    this.positions = [0, 0];
    this.currentPlayer = 0;
    this.gameOver = false;

    this.ladders = {
      2: 38, 7: 14, 8: 31, 15: 26, 21: 42,
      28: 84, 36: 44, 51: 67, 71: 91, 78: 98, 87: 94
    };

    this.snakes = {
      16: 6, 46: 25, 49: 11, 62: 19, 64: 60,
      74: 53, 89: 68, 92: 88, 95: 75, 99: 80
    };
  }

  play(die1, die2) {
    if (this.gameOver) return "Game over!";

    const player = this.currentPlayer;
    let pos = this.positions[player] + die1 + die2;

    // bounce back if overshooting 100
    if (pos > 100) {
      pos = 200 - pos;
    }

    // apply ladders and snakes
    if (this.ladders[pos]) pos = this.ladders[pos];
    else if (this.snakes[pos]) pos = this.snakes[pos];

    this.positions[player] = pos;

    if (pos === 100) {
      this.gameOver = true;
      return `Player ${player + 1} Wins!`;
    }

    // alternate unless doubles were rolled
    if (die1 !== die2) {
      this.currentPlayer = 1 - this.currentPlayer;
    }

    return `Player ${player + 1} is on square ${pos}`;
  }
}