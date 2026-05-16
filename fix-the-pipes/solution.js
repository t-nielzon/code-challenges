function fixThePipes(map, start, end) {
  const rows = map.length;
  const cols = map[0].length;
  const grid = map.map(r => r.split(''));

  const colsInfo = [];
  for (let c = 0; c < cols; c++) {
    let top = -1, bot = -1;
    for (let r = 0; r < rows; r++) {
      if (grid[r][c] === 'x') {
        if (top === -1) top = r;
        bot = r;
      }
    }
    colsInfo.push({ top, bot });
  }

  let entry = start;
  for (let c = 0; c < cols; c++) {
    const { top, bot } = colsInfo[c];
    let exit;
    if (top === bot) {
      exit = top;
      grid[top][c] = '━';
    } else if (entry === top) {
      exit = bot;
      grid[top][c] = '┓';
      for (let r = top + 1; r < bot; r++) grid[r][c] = '┃';
      grid[bot][c] = '┗';
    } else {
      exit = top;
      grid[bot][c] = '┛';
      for (let r = top + 1; r < bot; r++) grid[r][c] = '┃';
      grid[top][c] = '┏';
    }
    entry = exit;
  }

  return grid.map(r => r.join(''));
}