function countBlackSquares(width, height, resolution) {
  if (width === 0 || height === 0) return 0;
  
  const full_cols = Math.floor(width / resolution);
  const remainder_width = width % resolution;
  const full_rows = Math.floor(height / resolution);
  const remainder_height = height % resolution;
  
  const res_sq = resolution * resolution;
  let count = 0;
  
  const even_rows = Math.ceil(full_rows / 2);
  const odd_rows = Math.floor(full_rows / 2);
  const odd_cols = Math.floor(full_cols / 2);
  const even_cols = Math.ceil(full_cols / 2);
  
  count += even_rows * odd_cols * res_sq;
  count += odd_rows * even_cols * res_sq;
  
  if (remainder_width > 0) {
    if (full_cols % 2 === 1) {
      count += even_rows * remainder_width * resolution;
    } else {
      count += odd_rows * remainder_width * resolution;
    }
  }
  
  if (remainder_height > 0) {
    if (full_rows % 2 === 0) {
      count += odd_cols * resolution * remainder_height;
    } else {
      count += even_cols * resolution * remainder_height;
    }
  }
  
  if (remainder_width > 0 && remainder_height > 0) {
    if ((full_rows + full_cols) % 2 === 1) {
      count += remainder_width * remainder_height;
    }
  }
  
  return count;
}