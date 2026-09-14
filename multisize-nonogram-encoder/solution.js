function nonogramEncoder(nonogram) {
  const getClues = (sequence) => {
    const clues = [];
    let count = 0;
    for (let cell of sequence) {
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
    return clues;
  };

  const rowClues = nonogram.map(row => getClues(row));

  const columnClues = [];
  const cols = nonogram[0].length;
  for (let col = 0; col < cols; col++) {
    const column = nonogram.map(row => row[col]);
    columnClues.push(getClues(column));
  }

  return [columnClues, rowClues];
}