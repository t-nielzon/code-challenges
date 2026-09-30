function getSteppinOn(str) {
  const words = str.split(' ');
  
  // Calculate grid dimensions
  let maxRow = 0;
  let maxCol = 0;
  let row = 0;
  let col = 0;
  let direction = 'right';
  
  for (let i = 0; i < words.length; i++) {
    const word = words[i];
    
    if (direction === 'right') {
      maxCol = Math.max(maxCol, col + word.length - 1);
      col += word.length - 1;
    } else {
      maxRow = Math.max(maxRow, row + word.length - 1);
      row += word.length - 1;
    }
    
    direction = direction === 'right' ? 'down' : 'right';
  }
  
  // Create grid filled with spaces
  const grid = Array.from({ length: maxRow + 1 }, () => 
    Array.from({ length: maxCol + 1 }, () => ' ')
  );
  
  // Place words in the grid
  row = 0;
  col = 0;
  direction = 'right';
  
  for (let i = 0; i < words.length; i++) {
    const word = words[i];
    
    if (direction === 'right') {
      for (let j = 0; j < word.length; j++) {
        grid[row][col + j] = word[j];
      }
      col += word.length - 1;
    } else {
      for (let j = 0; j < word.length; j++) {
        grid[row + j][col] = word[j];
      }
      row += word.length - 1;
    }
    
    direction = direction === 'right' ? 'down' : 'right';
  }
  
  return grid;
}