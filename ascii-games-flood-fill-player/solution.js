function floodFillGame(board, moves) {
  let grid = board.split('\n').map(row => row.split(''));

  for (const { x, y, color } of moves) {
    const original = grid[y][x];
    if (original === String(color)) continue;
    const target = String(color);
    const rows = grid.length;
    const cols = grid[0].length;
    const stack = [[x, y]];
    grid[y][x] = target;
    while (stack.length) {
      const [cx, cy] = stack.pop();
      for (const [dx, dy] of [[1,0],[-1,0],[0,1],[0,-1]]) {
        const nx = cx + dx, ny = cy + dy;
        if (nx >= 0 && nx < cols && ny >= 0 && ny < rows && grid[ny][nx] === original) {
          grid[ny][nx] = target;
          stack.push([nx, ny]);
        }
      }
    }
  }

  return grid.map(row => row.join('')).join('\n');
}