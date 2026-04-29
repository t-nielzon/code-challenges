function peakHeight(mountain) {
  const rows = mountain.length;
  const cols = Math.max(...mountain.map(r => r.length));
  const grid = mountain.map(r => {
    const arr = [];
    for (let i = 0; i < cols; i++) arr.push(r[i] === '^' ? Infinity : 0);
    return arr;
  });

  let changed = true;
  let level = 0;
  while (changed) {
    changed = false;
    level++;
    const next = grid.map(r => r.slice());
    for (let i = 0; i < rows; i++) {
      for (let j = 0; j < cols; j++) {
        if (grid[i][j] >= level) {
          const up = i > 0 ? grid[i-1][j] : 0;
          const down = i < rows-1 ? grid[i+1][j] : 0;
          const left = j > 0 ? grid[i][j-1] : 0;
          const right = j < cols-1 ? grid[i][j+1] : 0;
          const minN = Math.min(up, down, left, right);
          if (minN >= level - 1 && minN < level) {
            // neighbor was peeled at previous level, this stays at level
          }
          if (Math.min(up, down, left, right) < level) {
            next[i][j] = Math.min(grid[i][j], Math.min(up, down, left, right) + 1);
          }
        }
      }
    }
    let max = 0;
    for (let i = 0; i < rows; i++) {
      for (let j = 0; j < cols; j++) {
        if (next[i][j] !== grid[i][j]) changed = true;
        grid[i][j] = next[i][j];
        if (grid[i][j] > max) max = grid[i][j];
      }
    }
    if (!changed) return max;
  }
  return level;
}