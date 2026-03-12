function squareSumsRow(n) {
  const isSquare = num => {
    const s = Math.sqrt(num);
    return s === Math.floor(s);
  };

  // build adjacency list: for each number 1..n, which other numbers can follow it
  const adj = Array.from({ length: n + 1 }, () => []);
  for (let i = 1; i <= n; i++) {
    for (let j = i + 1; j <= n; j++) {
      if (isSquare(i + j)) {
        adj[i].push(j);
        adj[j].push(i);
      }
    }
  }

  const path = [];
  const used = new Uint8Array(n + 1);

  function dfs() {
    if (path.length === n) return true;
    const last = path[path.length - 1];
    const neighbors = adj[last];
    for (let i = 0; i < neighbors.length; i++) {
      const next = neighbors[i];
      if (!used[next]) {
        used[next] = 1;
        path.push(next);
        if (dfs()) return true;
        path.pop();
        used[next] = 0;
      }
    }
    return false;
  }

  // try each starting number, prefer nodes with fewer connections first
  const starts = Array.from({ length: n }, (_, i) => i + 1);
  starts.sort((a, b) => adj[a].length - adj[b].length);

  // sort adjacency lists by degree (Warnsdorff-like heuristic)
  for (let i = 1; i <= n; i++) {
    adj[i].sort((a, b) => adj[a].length - adj[b].length);
  }

  for (const start of starts) {
    path.length = 0;
    used.fill(0);
    used[start] = 1;
    path.push(start);
    if (dfs()) return path;
  }

  return false;
}

module.exports = squareSumsRow;