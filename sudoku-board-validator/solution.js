function validSolution(board) {
  const isValidGroup = (nums) => {
    const sorted = [...nums].sort().join('');
    return sorted === '123456789';
  };

  for (let i = 0; i < 9; i++) {
    if (!isValidGroup(board[i])) return false;
    if (!isValidGroup(board.map(row => row[i]))) return false;
  }

  for (let r = 0; r < 9; r += 3) {
    for (let c = 0; c < 9; c += 3) {
      const block = [];
      for (let dr = 0; dr < 3; dr++) {
        for (let dc = 0; dc < 3; dc++) {
          block.push(board[r + dr][c + dc]);
        }
      }
      if (!isValidGroup(block)) return false;
    }
  }

  return true;
}