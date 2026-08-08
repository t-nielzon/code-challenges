function countBlack(width, height, resolution) {
  if (width === 0 || height === 0) return 0;
  
  const fullCols = Math.floor(width / resolution);
  const fullRows = Math.floor(height / resolution);
  const partialWidth = width % resolution;
  const partialHeight = height % resolution;
  
  let count = 0;
  
  // region 1: full resolution×resolution blocks arranged in a checkerboard
  if (fullCols > 0 && fullRows > 0) {
    const total = fullCols * fullRows;
    const numBlackBlocks = Math.floor(total / 2);
    count += numBlackBlocks * resolution * resolution;
  }
  
  // region 2: right edge strip (partial width, full height rows)
  if (partialWidth > 0 && fullRows > 0) {
    // first block at column fullCols; it's black if fullCols is odd
    const numBlackBlocks = (fullCols % 2 === 1) ? Math.ceil(fullRows / 2) : Math.floor(fullRows / 2);
    count += numBlackBlocks * partialWidth * resolution;
  }
  
  // region 3: bottom edge strip (full width columns, partial height)
  if (partialHeight > 0 && fullCols > 0) {
    // first block at row fullRows; it's black if fullRows is odd
    const numBlackBlocks = (fullRows % 2 === 1) ? Math.ceil(fullCols / 2) : Math.floor(fullCols / 2);
    count += numBlackBlocks * resolution * partialHeight;
  }
  
  // region 4: corner block (partial width and height)
  if (partialWidth > 0 && partialHeight > 0) {
    // block at (fullCols, fullRows) is black if sum of indices is odd
    if ((fullCols + fullRows) % 2 === 1) {
      count += partialWidth * partialHeight;
    }
  }
  
  return count;
}