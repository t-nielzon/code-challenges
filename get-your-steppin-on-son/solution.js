function stepIt(words) {
  const wordList = words.split(' ').filter(w => w.length > 0);
  if (wordList.length === 0) return [];
  
  // Calculate positions of each word
  const positions = [];
  let row = 0, col = 0;
  
  for (let i = 0; i < wordList.length; i++) {
    const word = wordList[i];
    const isHorizontal = i % 2 === 0;
    
    positions.push({
      word: word,
      row: row,
      col: col,
      isHorizontal: isHorizontal
    });
    
    // Move to next word's starting position
    if (isHorizontal) {
      col += word.length - 1;
    } else {
      row += word.length - 1;
    }
  }
  
  // Determine grid dimensions
  let maxRow = 0, maxCol = 0;
  for (let pos of positions) {
    if (pos.isHorizontal) {
      maxRow = Math.max(maxRow, pos.row);
      maxCol = Math.max(maxCol, pos.col + pos.word.length - 1);
    } else {
      maxRow = Math.max(maxRow, pos.row + pos.word.length - 1);
      maxCol = Math.max(maxCol, pos.col);
    }
  }
  
  // Create and initialize grid with spaces
  const grid = Array.from({ length: maxRow + 1 }, () => 
    Array(maxCol + 1).fill(' ')
  );
  
  // Place each word in the grid
  for (let pos of positions) {
    const word = pos.word;
    if (pos.isHorizontal) {
      for (let i = 0; i < word.length; i++) {
        grid[pos.row][pos.col + i] = word[i];
      }
    } else {
      for (let i = 0; i < word.length; i++) {
        grid[pos.row + i][pos.col] = word[i];
      }
    }
  }
  
  return grid;
}