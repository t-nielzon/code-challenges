function stepIt(str) {
  const words = str.split(' ');
  
  // Calculate grid dimensions
  let row = 0, col = 0;
  let maxRow = 0, maxCol = 0;
  let goingRight = true;
  
  for (let word of words) {
    if (goingRight) {
      maxCol = Math.max(maxCol, col + word.length - 1);
      col += word.length - 1;
      maxRow = Math.max(maxRow, row);
    } else {
      maxRow = Math.max(maxRow, row + word.length - 1);
      row += word.length - 1;
      maxCol = Math.max(maxCol, col);
    }
    goingRight = !goingRight;
  }
  
  // Create grid filled with spaces
  const grid = Array(maxRow + 1).fill(null).map(() => Array(maxCol + 1).fill(' '));
  
  // Place words in alternating directions (right, down, right, down, ...)
  row = 0, col = 0;
  goingRight = true;
  
  for (let word of words) {
    if (goingRight) {
      for (let i = 0; i < word.length; i++) {
        grid[row][col + i] = word[i];
      }
      col += word.length - 1;
    } else {
      for (let i = 0; i < word.length; i++) {
        grid[row + i][col] = word[i];
      }
      row += word.length - 1;
    }
    goingRight = !goingRight;
  }
  
  return grid;
}