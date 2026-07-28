function countBlackSquares(width, height, resolution) {
  if (width === 0 || height === 0) return 0;
  
  const full_cols = Math.floor(width / resolution);
  const full_rows = Math.floor(height / resolution);
  const partial_width = width % resolution;
  const partial_height = height % resolution;
  
  let count = 0;
  
  // count from full resolution x resolution grid
  const full_black = Math.floor((full_cols * full_rows) / 2);
  count += full_black * resolution * resolution;
  
  // count from right edge (if width not evenly divisible by resolution)
  if (partial_width > 0) {
    const col = full_cols;
    const black_rows = col % 2 === 0 ? Math.floor(full_rows / 2) : Math.ceil(full_rows / 2);
    count += black_rows * partial_width * resolution;
  }
  
  // count from bottom edge (if height not evenly divisible by resolution)
  if (partial_height > 0) {
    const row = full_rows;
    const black_cols = row % 2 === 0 ? Math.floor(full_cols / 2) : Math.ceil(full_cols / 2);
    count += black_cols * resolution * partial_height;
  }
  
  // count from bottom-right corner (if both edges are partial)
  if (partial_width > 0 && partial_height > 0) {
    const col = full_cols;
    const row = full_rows;
    if ((col + row) % 2 === 1) {
      count += partial_width * partial_height;
    }
  }
  
  return count;
}