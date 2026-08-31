function omnifinite(n) {
  const cells = [];
  
  // A horizontal line of 2*n cells dies out after exactly n generations:
  // Gen 0: cells at [0,0], [1,0], ..., [2n-1,0]
  // Each generation, the line shrinks from both ends as edge cells die (only 1 neighbor)
  // Inner cells survive with 2 neighbors until only 2 remain
  // Gen n: empty
  
  for (let i = 0; i < 2 * n; i++) {
    cells.push([i, 0]);
  }
  
  return cells;
}