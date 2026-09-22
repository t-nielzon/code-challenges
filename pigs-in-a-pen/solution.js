class Game {
  constructor(board) {
    this.board = board;
  }

  play(lines) {
    const n = this.board;
    const filled = new Set(lines);
    
    // h_lines: (n+1) * n lines numbered 1 to (n+1)*n
    // v_lines: n * (n+1) lines numbered (n+1)*n + 1 to total
    const h_count = (n + 1) * n;
    const v_offset = h_count;
    
    const getSides = (r, c) => {
      const top = r * n + c + 1;
      const bottom = (r + 1) * n + c + 1;
      const left = v_offset + r * (n + 1) + c + 1;
      const right = v_offset + r * (n + 1) + (c + 1) + 1;
      return [top, bottom, left, right];
    };
    
    let changed = true;
    while (changed) {
      changed = false;
      
      for (let r = 0; r < n; r++) {
        for (let c = 0; c < n; c++) {
          const sides = getSides(r, c);
          const present = sides.filter(s => filled.has(s));
          
          if (present.length === 3) {
            for (const side of sides) {
              if (!filled.has(side)) {
                filled.add(side);
                changed = true;
              }
            }
          }
        }
      }
    }
    
    return Array.from(filled).sort((a, b) => a - b);
  }
}