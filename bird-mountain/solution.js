function peakHeight(mountain) {
  const grid = mountain.map(row => row.split(''));
  const rows = grid.length;
  const isHat = (i, j) =>
    i >= 0 && i < rows && j >= 0 && j < grid[i].length && grid[i][j] === '^';

  let height = 0;
  while (true) {
    const toMark = [];
    let anyHat = false;
    for (let i = 0; i < rows; i++) {
      for (let j = 0; j < grid[i].length; j++) {
        if (grid[i][j] !== '^') continue;
        anyHat = true;
        if (!isHat(i - 1, j) || !isHat(i + 1, j) ||
            !isHat(i, j - 1) || !isHat(i, j + 1)) {
          toMark.push([i, j]);
        }
      }
    }
    if (!anyHat) break;
    height++;
    for (const [i, j] of toMark) grid[i][j] = String(height % 10);
  }
  return height;
}