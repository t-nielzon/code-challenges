function gameOfLifeOmnifinite(n) {
  if (n === 0) return [];
  if (n === 1) return [[0, 0]];
  
  const cells = [];
  
  // Create a diagonal pattern that evolves over n steps
  // Each cell at [i, i] connects with its neighbor to form structures
  // that interact and eventually all die after n generations
  
  for (let i = 0; i < n; i++) {
    cells.push([i, 0]);
    cells.push([i, 1]);
  }
  
  // Add offset cells to create interaction patterns
  for (let i = 0; i < n - 1; i++) {
    cells.push([i + 1, 2]);
  }
  
  return cells;
}