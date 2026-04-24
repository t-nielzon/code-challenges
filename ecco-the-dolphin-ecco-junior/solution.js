function findPath(maze) {
  const grid = maze.split('\n');
  const h = grid.length;
  if (h === 0) return null;
  const w = grid[0].length;

  let ecco = null, tara = null;
  const sharks = [], jellies = [], urchins = [];

  for (let y = 0; y < h; y++) {
    for (let x = 0; x < w; x++) {
      const c = grid[y][x];
      if (c === 'e') ecco = [y, x];
      else if (c === 't') tara = [y, x];
      else if (c === 'S') sharks.push([y, x]);
      else if (c === 'Y') jellies.push([y, x]);
      else if (c === 'O') urchins.push([y, x]);
    }
  }

  if (!ecco || !tara) return null;

  const danger = Array.from({ length: h }, () => new Array(w).fill(false));

  for (const [sy, sx] of sharks) {
    let left = sx;
    while (left - 1 >= 0 && grid[sy][left - 1] !== '#') left--;
    let right = sx;
    while (right + 1 < w && grid[sy][right + 1] !== '#') right++;
    for (let x = left; x <= right; x++) danger[sy][x] = true;
  }

  for (const [jy, jx] of jellies) {
    let top = jy;
    while (top - 1 >= 0 && grid[top - 1][jx] !== '#' && grid[top - 1][jx] !== '.') top--;
    let bottom = jy;
    while (bottom + 1 < h && grid[bottom + 1][jx] !== '#' && grid[bottom + 1][jx] !== '.') bottom++;
    for (let y = top; y <= bottom; y++) danger[y][jx] = true;
  }

  for (const [uy, ux] of urchins) {
    for (let dy = -1; dy <= 1; dy++) {
      for (let dx = -1; dx <= 1; dx++) {
        const ny = uy + dy, nx = ux + dx;
        if (ny >= 0 && ny < h && nx >= 0 && nx < w) {
          danger[ny][nx] = true;
        }
      }
    }
  }

  danger[ecco[0]][ecco[1]] = false;
  danger[tara[0]][tara[1]] = false;

  const passable = (y, x) => {
    if (y < 0 || y >= h || x < 0 || x >= w) return false;
    const c = grid[y][x];
    if (c === '#' || c === '.') return false;
    if (danger[y][x]) return false;
    return true;
  };

  const visited = Array.from({ length: h }, () => new Array(w).fill(false));
  const parent = Array.from({ length: h }, () => new Array(w).fill(null));
  const queue = [ecco];
  let head = 0;
  visited[ecco[0]][ecco[1]] = true;

  const dirs = [[-1, 0], [1, 0], [0, -1], [0, 1]];

  while (head < queue.length) {
    const [y, x] = queue[head++];
    if (y === tara[0] && x === tara[1]) {
      const path = [];
      let cur = [y, x];
      while (cur) {
        path.push(cur);
        cur = parent[cur[0]][cur[1]];
      }
      path.reverse();
      return path;
    }
    for (const [dy, dx] of dirs) {
      const ny = y + dy, nx = x + dx;
      if (!visited[ny] || visited[ny][nx]) {
        if (ny >= 0 && ny < h && nx >= 0 && nx < w && visited[ny][nx]) continue;
      }
      if (passable(ny, nx) && !visited[ny][nx]) {
        visited[ny][nx] = true;
        parent[ny][nx] = [y, x];
        queue.push([ny, nx]);
      }
    }
  }

  return null;
}