function findPath(maze) {
  const grid = maze.split('\n').map(row => row.split(''));
  const H = grid.length;
  const W = grid[0].length;

  const k = (y, x) => y + ',' + x;
  const danger = new Set();
  let start = null, goal = null;

  for (let y = 0; y < H; y++) {
    for (let x = 0; x < (grid[y].length); x++) {
      const c = grid[y][x];

      if (c === 'e') {
        start = [y, x];
      } else if (c === 't') {
        goal = [y, x];
      } else if (c === 'S') {
        // sharks sweep their row, bouncing only off rocks
        for (let xx = x; xx < W && grid[y][xx] !== '#'; xx++) danger.add(k(y, xx));
        for (let xx = x; xx >= 0 && grid[y][xx] !== '#'; xx--) danger.add(k(y, xx));
      } else if (c === 'Y') {
        // jellyfish sweep their column, bouncing off rocks and air
        for (let yy = y; yy < H && grid[yy][x] !== '#' && grid[yy][x] !== '.'; yy++) danger.add(k(yy, x));
        for (let yy = y; yy >= 0 && grid[yy][x] !== '#' && grid[yy][x] !== '.'; yy--) danger.add(k(yy, x));
      } else if (c === 'O') {
        // urchins expand into every adjacent tile, diagonals included
        for (let dy = -1; dy <= 1; dy++) {
          for (let dx = -1; dx <= 1; dx++) {
            const ny = y + dy, nx = x + dx;
            if (ny >= 0 && ny < H && nx >= 0 && nx < W) danger.add(k(ny, nx));
          }
        }
      }
    }
  }

  if (!start || !goal) return null;

  const passable = (y, x) => {
    if (y < 0 || y >= H || x < 0 || x >= W) return false;
    if (danger.has(k(y, x))) return false;
    const c = grid[y][x];
    return c === '~' || c === 'e' || c === 't';
  };

  const dirs = [[-1, 0], [1, 0], [0, -1], [0, 1]];
  const prev = new Map();
  prev.set(k(start[0], start[1]), null);
  const queue = [start];

  while (queue.length) {
    const [y, x] = queue.shift();

    if (y === goal[0] && x === goal[1]) {
      const path = [];
      let cy = y, cx = x;
      let cur = k(cy, cx);
      while (cur !== null) {
        path.push([cy, cx]);
        const p = prev.get(cur);
        if (p === null) break;
        cy = p[0];
        cx = p[1];
        cur = k(cy, cx);
      }
      path.reverse();
      return path;
    }

    for (const [dy, dx] of dirs) {
      const ny = y + dy, nx = x + dx;
      const nk = k(ny, nx);
      if (passable(ny, nx) && !prev.has(nk)) {
        prev.set(nk, [y, x]);
        queue.push([ny, nx]);
      }
    }
  }

  return null;
}