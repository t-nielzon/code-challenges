class Game {
  constructor(board) {
    this.board = board;
  }

  play(lines) {
    const filled = new Set(lines);
    let changed = true;

    while (changed) {
      changed = false;
      
      for (let brow = 0; brow < this.board; brow++) {
        for (let bcol = 0; bcol < this.board; bcol++) {
          const [top, right, bottom, left] = this.getBoxEdges(brow, bcol);
          const edges = [top, right, bottom, left];
          
          const filledCount = edges.filter(e => filled.has(e)).length;
          
          if (filledCount === 3) {
            const missing = edges.find(e => !filled.has(e));
            filled.add(missing);
            changed = true;
          }
        }
      }
    }

    return Array.from(filled).sort((a, b) => a - b);
  }

  getBoxEdges(brow, bcol) {
    const n = this.board;
    // horizontal lines: 1 to n(n+1), numbered by row then column
    // vertical lines: n(n+1)+1 to 2n(n+1), numbered by row then column
    const top = 1 + brow * n + bcol;
    const bottom = 1 + (brow + 1) * n + bcol;
    const left = 1 + n * (n + 1) + brow * (n + 1) + bcol;
    const right = 1 + n * (n + 1) + brow * (n + 1) + (bcol + 1);
    
    return [top, right, bottom, left];
  }
}