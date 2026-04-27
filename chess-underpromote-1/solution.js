function chessPawn(start, target) {
  if (start === target) return 0;
  const sf = start.charCodeAt(0) - 97, sr = +start[1];
  const tf = target.charCodeAt(0) - 97, tr = +target[1];

  let best = Infinity;

  if (sf === tf && tr >= sr) {
    best = Math.min(best, tr - sr);
  }

  const pushes = 8 - sr;
  const qd = queenDist(sf, 8, tf, tr);
  const kd = knightDist(sf, 8, tf, tr);
  best = Math.min(best, pushes + Math.min(qd, kd));

  return best;
}

function queenDist(x1, y1, x2, y2) {
  if (x1 === x2 && y1 === y2) return 0;
  if (x1 === x2 || y1 === y2 || Math.abs(x1 - x2) === Math.abs(y1 - y2)) return 1;
  return 2;
}

function knightDist(x1, y1, x2, y2) {
  const moves = [[1,2],[2,1],[-1,2],[-2,1],[1,-2],[2,-1],[-1,-2],[-2,-1]];
  const visited = new Set();
  const q = [[x1, y1, 0]];
  visited.add(x1 * 10 + y1);
  while (q.length) {
    const [x, y, d] = q.shift();
    if (x === x2 && y === y2) return d;
    for (const [dx, dy] of moves) {
      const nx = x + dx, ny = y + dy;
      if (nx < 0 || nx > 7 || ny < 1 || ny > 8) continue;
      const k = nx * 10 + ny;
      if (visited.has(k)) continue;
      visited.add(k);
      q.push([nx, ny, d + 1]);
    }
  }
  return Infinity;
}