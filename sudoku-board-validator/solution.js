function validSolution(board) {
  const isValidGroup = (nums) => {
    const sorted = [...nums].sort().join('');
    return sorted === '123456789';
  };

  for (let i = 0; i < 9; i++) {
    // check row
    if (!isValidGroup(board[i])) return false;
    // check column
    if (!isValidGroup(board.map(row => row[i]))) return false;
  }

  // check 3x3 blocks
  for (let r = 0; r < 9; r += 3) {
    for (let c = 0; c < 9; c += 3) {
      const block = [];
      for (let i = 0; i < 3; i++) {
        for (let j = 0; j < 3; j++) {
          block.push(board[r + i][c + j]);
        }
      }
      if (!isValidGroup(block)) return false;
    }
  }

  return true;
}