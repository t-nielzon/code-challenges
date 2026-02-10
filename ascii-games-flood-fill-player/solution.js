function floodFillPlayer(board, moves) {
  let grid = board.split('\n').map(row => row.split(''));

  for (const move of moves) {
    const { x, y, color } = move;
    const newColor = String(color);
    const oldColor = grid[y][x];
    if (oldColor === newColor) continue;

    const rows = grid.length;
    const cols = grid[0].length;
    const stack = [[x, y]];
    grid[y][x] = newColor;

    while (stack.length) {
      const [cx, cy] = stack.pop();
      for (const [dx, dy] of [[1, 0], [-1, 0], [0, 1], [0, -1]]) {
        const nx = cx + dx;
        const ny = cy + dy;
        if (nx >= 0 && nx < cols && ny >= 0 && ny < rows && grid[ny][nx] === oldColor) {
          grid[ny][nx] = newColor;
          stack.push([nx, ny]);
        }
      }
    }
  }

  return grid.map(row => row.join('')).join('\n');
}

module.exports = floodFillPlayer;