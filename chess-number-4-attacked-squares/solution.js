// Solution: solution.js
function attackedSquares(board) {
  const result = Array(8).fill(null).map(() => Array(8).fill("  "));
  
  const isValid = (r, c) => r >= 0 && r < 8 && c >= 0 && c < 8;
  
  const markAttacked = (r, c, color) => {
    if (!isValid(r, c)) return;
    if (color === 'white') {
      if (result[r][c] === "  ") result[r][c] = "W ";
      else if (result[r][c] === " B") result[r][c] = "WB";
    } else {
      if (result[r][c] === "  ") result[r][c] = " B";
      else if (result[r][c] === "W ") result[r][c] = "WB";
    }
  };
  
  const getAttackedSquares = (piece, r, c) => {
    const attacks = [];
    const type = piece.toLowerCase();
    const isWhite = piece === piece.toUpperCase();
    
    if (type === 'p') {
      if (isWhite) {
        if (isValid(r - 1, c - 1)) attacks.push([r - 1, c - 1]);
        if (isValid(r - 1, c + 1)) attacks.push([r - 1, c + 1]);
      } else {
        if (isValid(r + 1, c - 1)) attacks.push([r + 1, c - 1]);
        if (isValid(r + 1, c + 1)) attacks.push([r + 1, c + 1]);
      }
    } else if (type === 'n') {
      const moves = [[-2, -1], [-2, 1], [-1, -2], [-1, 2], [1, -2], [1, 2], [2, -1], [2, 1]];
      for (const [dr, dc] of moves) {
        const nr = r + dr, nc = c + dc;
        if (isValid(nr, nc)) attacks.push([nr, nc]);
      }
    } else if (type === 'b') {
      const directions = [[-1, -1], [-1, 1], [1, -1], [1, 1]];
      for (const [dr, dc] of directions) {
        for (let i = 1; i < 8; i++) {
          const nr = r + dr * i, nc = c + dc * i;
          if (!isValid(nr, nc)) break;
          attacks.push([nr, nc]);
          if (board[nr][nc] !== ' ') break;
        }
      }
    } else if (type === 'r') {
      const directions = [[-1, 0], [1, 0], [0, -1], [0, 1]];
      for (const [dr, dc] of directions) {
        for (let i = 1; i < 8; i++) {
          const nr = r + dr * i, nc = c + dc * i;
          if (!isValid(nr, nc)) break;
          attacks.push([nr, nc]);
          if (board[nr][nc] !== ' ') break;
        }
      }
    } else if (type === 'q') {
      const directions = [[-1, -1], [-1, 0], [-1, 1], [0, -1], [0, 1], [1, -1], [1, 0], [1, 1]];
      for (const [dr, dc] of directions) {
        for (let i = 1; i < 8; i++) {
          const nr = r + dr * i, nc = c + dc * i;
          if (!isValid(nr, nc)) break;
          attacks.push([nr, nc]);
          if (board[nr][nc] !== ' ') break;
        }
      }
    } else if (type === 'k') {
      const directions = [[-1, -1], [-1, 0], [-1, 1], [0, -1], [0, 1], [1, -1], [1, 0], [1, 1]];
      for (const [dr, dc] of directions) {
        const nr = r + dr, nc = c + dc;
        if (isValid(nr, nc)) attacks.push([nr, nc]);
      }
    }
    
    return attacks;
  };
  
  for (let r = 0; r < 8; r++) {
    for (let c = 0; c < 8; c++) {
      const piece = board[r][c];
      if (piece === ' ') continue;
      
      const isWhite = piece === piece.toUpperCase();
      const attacks = getAttackedSquares(piece, r, c);
      
      for (const [ar, ac] of attacks) {
        markAttacked(ar, ac, isWhite ? 'white' : 'black');
      }
    }
  }
  
  return result;
}