class Game {
  constructor(board) {
    this.board = board;
  }
  
  play(lines) {
    let result = new Set(lines);
    let changed = true;
    
    while (changed) {
      changed = false;
      
      for (let row = 0; row < this.board; row++) {
        for (let col = 0; col < this.board; col++) {
          const squareLines = this.getSquareLines(row, col);
          const present = squareLines.filter(line => result.has(line));
          
          if (present.length === 3) {
            const missing = squareLines.find(line => !result.has(line));
            result.add(missing);
            changed = true;
          }
        }
      }
    }
    
    return Array.from(result).sort((a, b) => a - b);
  }
  
  getSquareLines(row, col) {
    const levelSize = 2 * this.board + 1;
    const top = 1 + row * levelSize + col;
    const bottom = 1 + (row + 1) * levelSize + col;
    const left = this.board + 1 + row * levelSize + col;
    const right = this.board + 1 + row * levelSize + col + 1;
    
    return [top, bottom, left, right];
  }
}