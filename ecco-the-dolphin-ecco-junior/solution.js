function solve(maze) {
  const grid = maze.split('\n').map(r => r.split(''));
  const rows = grid.length;
  const cols = grid[0].length;

  let start, end;
  const sharks = [], jellyfish = [], urchins = [];

  for (let y = 0; y < rows; y++) {
    for (let x = 0; x < cols; x++) {
      const c = grid[y][x];
      if (c === 'e') start = [y, x];
      else if (c === 't') end = [y, x];
      else if (c === 'S') sharks.push([y, x]);
      else if (c === 'Y') jellyfish.push([y, x]);
      else if (c === 'O') urchins.push([y, x]);
    }
  }

  const dangerous = Array.from({ length: rows }, () => Array(cols).fill(false));

  // sharks patrol horizontally between rocks
  for (const [sy, sx] of sharks) {
    dangerous[sy][sx] = true;
    for (let x = sx - 1; x >= 0 && grid[sy][x] !== '#'; x--) dangerous[sy][x] = true;
    for (let x = sx + 1; x < cols && grid[sy][x] !== '#'; x++) dangerous[sy][x] = true;
  }

  // jellyfish patrol vertically between rocks and air
  for (const [jy, jx] of jellyfish) {
    dangerous[jy][jx] = true;
    for (let y = jy - 1; y >= 0 && grid[y][jx] !== '#' && grid[y][jx] !== '.'; y--) dangerous[y][jx] = true;
    for (let y = jy + 1; y < rows && grid[y][jx] !== '#' && grid[y][jx] !== '.'; y++) dangerous[y][jx] = true;
  }

  // urchins expand to all 8 adjacent tiles
  for (const [uy, ux] of urchins) {
    for (let dy = -1; dy <= 1; dy++) {
      for (let dx = -1; dx <= 1; dx++) {
        const ny = uy + dy, nx = ux + dx;
        if (ny >= 0 && ny < rows && nx >= 0 && nx < cols) dangerous[ny][nx] = true;
      }
    }
  }

  const isPassable = (y, x) =>
    y >= 0 && y < rows && x >= 0 && x < cols &&
    grid[y][x] !== '#' && grid[y][x] !== '.' &&
    !dangerous[y][x];

  // bfs
  const visited = Array.from({ length: rows }, () => Array(cols).fill(false));
  const parent = Array.from({ length: rows }, () => Array(cols).fill(null));
  const queue = [start];
  visited[start[0]][start[1]] = true;

  const dirs = [[0, 1], [0, -1], [1, 0], [-1, 0]];
  let found = false;

  while (queue.length > 0) {
    const [cy, cx] = queue.shift();
    if (cy === end[0] && cx === end[1]) { found = true; break; }
    for (const [dy, dx] of dirs) {
      const ny = cy + dy, nx = cx + dx;
      if (isPassable(ny, nx) && !visited[ny][nx]) {
        visited[ny][nx] = true;
        parent[ny][nx] = [cy, cx];
        queue.push([ny, nx]);
      }
    }
  }

  if (!found) return null;

  const path = [];
  let cur = end;
  while (cur) {
    path.push(cur);
    cur = parent[cur[0]][cur[1]];
  }
  path.reverse();
  return path;
}

module.exports = { solve };