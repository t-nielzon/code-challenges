function fixThePipes(map, start, end) {
  const isStr = typeof map[0] === 'string';
  const grid = map.map(r => isStr ? r.split('') : r.slice());
  const rows = grid.length;
  const cols = grid[0].length;

  let enterRow = start;
  for (let c = 0; c < cols; c++) {
    const xRows = [];
    for (let r = 0; r < rows; r++) {
      if (grid[r][c] === 'x') xRows.push(r);
    }
    if (xRows.length === 0) continue;

    const minR = xRows[0];
    const maxR = xRows[xRows.length - 1];

    let exitRow;
    if (c === cols - 1) {
      exitRow = end;
    } else {
      exitRow = (enterRow === minR) ? maxR : minR;
    }

    if (minR === maxR) {
      grid[minR][c] = '━';
    } else {
      grid[minR][c] = (minR === enterRow) ? '┓' : '┏';
      grid[maxR][c] = (maxR === enterRow) ? '┛' : '┗';
      for (let r = minR + 1; r < maxR; r++) {
        grid[r][c] = '┃';
      }
    }

    enterRow = exitRow;
  }

  return isStr ? grid.map(r => r.join('')) : grid;
}