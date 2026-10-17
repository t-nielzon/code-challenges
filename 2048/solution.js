function updateBoard(board, direction) {
  const copy = board.map(row => [...row]);
  const original = board.map(row => [...row]);
  
  if (direction === 0) { // up
    for (let col = 0; col < 4; col++) {
      let column = [copy[0][col], copy[1][col], copy[2][col], copy[3][col]];
      column = slideAndCombine(column);
      for (let row = 0; row < 4; row++) {
        copy[row][col] = column[row];
      }
    }
  } else if (direction === 1) { // left
    for (let row = 0; row < 4; row++) {
      copy[row] = slideAndCombine(copy[row]);
    }
  } else if (direction === 2) { // right
    for (let row = 0; row < 4; row++) {
      copy[row] = copy[row].reverse();
      copy[row] = slideAndCombine(copy[row]);
      copy[row] = copy[row].reverse();
    }
  } else if (direction === 3) { // down
    for (let col = 0; col < 4; col++) {
      let column = [copy[3][col], copy[2][col], copy[1][col], copy[0][col]];
      column = slideAndCombine(column);
      column.reverse();
      for (let row = 0; row < 4; row++) {
        copy[row][col] = column[row];
      }
    }
  }
  
  if (JSON.stringify(copy) === JSON.stringify(original)) {
    return original;
  }
  
  return copy;
}

function slideAndCombine(line) {
  let tiles = line.filter(x => x !== 0);
  
  let combined = [];
  let i = 0;
  while (i < tiles.length) {
    if (i + 1 < tiles.length && tiles[i] === tiles[i + 1]) {
      combined.push(tiles[i] + tiles[i + 1]);
      i += 2;
    } else {
      combined.push(tiles[i]);
      i += 1;
    }
  }
  
  while (combined.length < 4) {
    combined.push(0);
  }
  
  return combined;
}