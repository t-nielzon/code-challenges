function encodeNonogram(nonogram) {
  const rowClues = [];
  const columnClues = [];
  
  // Process rows
  for (let row of nonogram) {
    const clues = [];
    let count = 0;
    for (let cell of row) {
      if (cell === 1) {
        count++;
      } else {
        if (count > 0) {
          clues.push(count);
          count = 0;
        }
      }
    }
    if (count > 0) {
      clues.push(count);
    }
    rowClues.push(clues);
  }
  
  // Process columns
  const cols = nonogram[0].length;
  for (let colIdx = 0; colIdx < cols; colIdx++) {
    const clues = [];
    let count = 0;
    for (let rowIdx = 0; rowIdx < nonogram.length; rowIdx++) {
      const cell = nonogram[rowIdx][colIdx];
      if (cell === 1) {
        count++;
      } else {
        if (count > 0) {
          clues.push(count);
          count = 0;
        }
      }
    }
    if (count > 0) {
      clues.push(count);
    }
    columnClues.push(clues);
  }
  
  return [columnClues, rowClues];
}