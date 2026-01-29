function squares(x, y) {
  let total = 0;
  
  // For each possible square size from 1 to the smaller dimension
  for (let size = 1; size <= y; size++) {
    // Count how many squares of this size fit in the rectangle
    // Horizontally: (x - size + 1) positions
    // Vertically: (y - size + 1) positions
    total += (x - size + 1) * (y - size + 1);
  }
  
  return total;
}