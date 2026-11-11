/*
 * Pigs in a Pen (Dots and Boxes Game Solver)
 * 
 * Given a board size and a list of already-drawn lines, find all squares
 * that can be completed (where 3 out of 4 sides are already drawn).
 * A square can only be completed if exactly 3 of its 4 sides are present.
 * Return all lines including the original ones, sorted numerically with no duplicates.
 */

class Game {
  constructor(board) {
    this.board = board;
  }

  play(lines) {
    const completed = new Set(lines);
    let changed = true;
    
    while (changed) {
      changed = false;
      for (let i = 0; i < this.board; i++) {
        for (let j = 0; j < this.board; j++) {
          const sides = this.getSidesOfSquare(i, j);
          const presentCount = sides.filter(s => completed.has(s)).length;
          
          if (presentCount === 3) {
            const missingSide = sides.find(s => !completed.has(s));
            if (missingSide && !completed.has(missingSide)) {
              completed.add(missingSide);
              changed = true;
            }
          }
        }
      }
    }
    
    return Array.from(completed).sort((a, b) => a - b);
  }

  getSidesOfSquare(i, j) {
    const n = this.board;
    const total = n * (n + 1);
    
    const top = i * (n + 1) + j + 1;
    const bottom = (i + 1) * (n + 1) + j + 1;
    const left = total + i * (n + 1) + j + 1;
    const right = total + i * (n + 1) + (j + 1) + 1;
    
    return [top, bottom, left, right];
  }
}