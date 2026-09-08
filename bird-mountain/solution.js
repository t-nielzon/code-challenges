function mountainHeight(mountain) {
  if (!mountain || mountain.length === 0) return 0;
  
  // Parse mountain to find all caret positions
  let carets = new Set();
  mountain.forEach((line, row) => {
    for (let col = 0; col < line.length; col++) {
      if (line[col] === '^') {
        carets.add(`${row},${col}`);
      }
    }
  });
  
  if (carets.size === 0) return 0;
  
  let height = 0;
  
  while (carets.size > 0) {
    let edges = new Set();
    
    // Find all edge carets (those without neighbors in all 4 directions)
    for (let pos of carets) {
      let [row, col] = pos.split(',').map(Number);
      
      let hasUp = carets.has(`${row - 1},${col}`);
      let hasDown = carets.has(`${row + 1},${col}`);
      let hasLeft = carets.has(`${row},${col - 1}`);
      let hasRight = carets.has(`${row},${col + 1}`);
      
      if (!hasUp || !hasDown || !hasLeft || !hasRight) {
        edges.add(pos);
      }
    }
    
    if (edges.size === 0) break;
    
    for (let edge of edges) {
      carets.delete(edge);
    }
    
    height++;
  }
  
  return height;
}