function squareSumsRow(n) {
  // Precompute all perfect squares up to 2*n (max possible sum is n + (n-1))
  const maxSum = 2 * n;
  const squares = new Set();
  for (let i = 2; i * i <= maxSum; i++) {
    squares.add(i * i);
  }
  
  // Build adjacency list: for each number, find all valid neighbors
  const neighbors = new Array(n + 1).fill(null).map(() => []);
  for (let i = 1; i <= n; i++) {
    for (let j = i + 1; j <= n; j++) {
      if (squares.has(i + j)) {
        neighbors[i].push(j);
        neighbors[j].push(i);
      }
    }
  }
  
  // Sort neighbors by degree (fewest connections first) for better pruning
  for (let i = 1; i <= n; i++) {
    neighbors[i].sort((a, b) => neighbors[a].length - neighbors[b].length);
  }
  
  const used = new Array(n + 1).fill(false);
  const path = [];
  
  function backtrack() {
    if (path.length === n) {
      return true;
    }
    
    const last = path[path.length - 1];
    for (const next of neighbors[last]) {
      if (!used[next]) {
        used[next] = true;
        path.push(next);
        if (backtrack()) {
          return true;
        }
        path.pop();
        used[next] = false;
      }
    }
    return false;
  }
  
  // Try starting from each number, prioritizing those with fewer neighbors
  const startOrder = [];
  for (let i = 1; i <= n; i++) {
    startOrder.push(i);
  }
  startOrder.sort((a, b) => neighbors[a].length - neighbors[b].length);
  
  for (const start of startOrder) {
    used[start] = true;
    path.push(start);
    if (backtrack()) {
      return path;
    }
    path.pop();
    used[start] = false;
  }
  
  return false;
}

module.exports = { squareSumsRow };