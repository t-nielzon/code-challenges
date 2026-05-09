class SnakesLadders {
  constructor() {
    this.positions = [0, 0];
    this.current = 0;
    this.winner = null;
    this.ladders = {
      1: 38, 4: 14, 9: 31, 21: 42, 28: 84,
      36: 44, 51: 67, 71: 91, 80: 100
    };
    this.snakes = {
      16: 6, 47: 26, 49: 11, 56: 53, 62: 19,
      64: 60, 87: 24, 93: 73, 95: 75, 98: 78
    };
  }
  
  play(die1, die2) {
    if (this.winner !== null) return "Game over!";
    
    const player = this.current;
    let pos = this.positions[player] + die1 + die2;
    if (pos > 100) pos = 200 - pos;
    if (this.ladders[pos] !== undefined) pos = this.ladders[pos];
    else if (this.snakes[pos] !== undefined) pos = this.snakes[pos];
    
    this.positions[player] = pos;
    
    if (pos === 100) {
      this.winner = player + 1;
      return `Player ${player + 1} Wins!`;
    }
    
    const message = `Player ${player + 1} is on square ${pos}`;
    
    if (die1 !== die2) this.current = 1 - this.current;
    
    return message;
  }
}