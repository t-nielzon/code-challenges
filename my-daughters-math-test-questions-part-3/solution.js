function neighbor(s) {
  const lines = s.split('\n');
  const grid = [];
  const results = [];

  for (const line of lines) {
    const m = line.match(/\|\s*([^|_]+?)\s*\|\s*([^|_]+?)\s*\|\s*([^|_]+?)\s*\|/);
    if (m) grid.push([m[1].trim(), m[2].trim(), m[3].trim()]);
  }

  for (const line of lines) {
    const q = line.match(/(.+?):\s*My\s+(left|right|upstairs|downstairs)\s+neighbor\s+is\s+\(\?\)/);
    if (!q) continue;

    const name = q[1].trim();
    const dir = q[2];
    let row = -1, col = -1;

    for (let r = 0; r < 3; r++)
      for (let c = 0; c < 3; c++)
        if (grid[r][c] === name) { row = r; col = c; }

    let ans = 'nobody';
    if (dir === 'left' && col > 0) ans = grid[row][col - 1];
    else if (dir === 'right' && col < 2) ans = grid[row][col + 1];
    else if (dir === 'upstairs' && row > 0) ans = grid[row - 1][col];
    else if (dir === 'downstairs' && row < 2) ans = grid[row + 1][col];

    results.push(ans);
  }

  return results.join('\n');
}