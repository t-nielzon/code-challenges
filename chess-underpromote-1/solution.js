function underpromote(start, target) {
  if (start === target) return 0;

  const sf = start.charCodeAt(0) - 97;
  const sr = parseInt(start[1], 10) - 1;
  const tf = target.charCodeAt(0) - 97;
  const tr = parseInt(target[1], 10) - 1;

  if (sf === tf && tr >= sr) {
    return tr - sr;
  }

  const movesToRank8 = 7 - sr;
  const pf = sf, pr = 7;

  let queen;
  if (pf === tf && pr === tr) queen = 0;
  else if (pf === tf || pr === tr || Math.abs(pf - tf) === Math.abs(pr - tr)) queen = 1;
  else queen = 2;

  const knight = knightDist(pf, pr, tf, tr);

  return movesToRank8 + Math.min(queen, knight);
}

function knightDist(sf, sr, tf, tr) {
  if (sf === tf && sr === tr) return 0;
  const visited = Array.from({ length: 8 }, () => Array(8).fill(false));
  const queue = [[sf, sr, 0]];
  visited[sf][sr] = true;
  const deltas = [[1, 2], [2, 1], [-1, 2], [-2, 1], [1, -2], [2, -1], [-1, -2], [-2, -1]];
  while (queue.length) {
    const [f, r, d] = queue.shift();
    for (const [df, dr] of deltas) {
      const nf = f + df, nr = r + dr;
      if (nf >= 0 && nf < 8 && nr >= 0 && nr < 8 && !visited[nf][nr]) {
        if (nf === tf && nr === tr) return d + 1;
        visited[nf][nr] = true;
        queue.push([nf, nr, d + 1]);
      }
    }
  }
  return Infinity;
}