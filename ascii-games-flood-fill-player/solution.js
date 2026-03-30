function floodFillPlayer(board, moves) {
  let grid = board.split('\n').map(row => row.split(''));

  for (const { x, y, color } of moves) {
    const original = grid[y][x];
    if (original === String(color)) continue;
    const stack = [[x, y]];
    const visited = new Set();
    while (stack.length) {
      const [cx, cy] = stack.pop();
      const key = cy * grid[0].length + cx;
      if (visited.has(key)) continue;
      if (cx < 0 || cy < 0 || cy >= grid.length || cx >= grid[0].length) continue;
      if (grid[cy][cx] !== original) continue;
      visited.add(key);
      grid[cy][cx] = String(color);
      stack.push([cx + 1, cy], [cx - 1, cy], [cx, cy + 1], [cx, cy - 1]);
    }
  }

  return grid.map(row => row.join('')).join('\n');
}