function checkerboard(width, height, resolution) {
  const fullBlocksX = Math.floor(width / resolution);
  const fullBlocksY = Math.floor(height / resolution);
  const partialWidthX = width % resolution;
  const partialHeightY = height % resolution;
  
  let count = 0;
  
  // Count black squares in the grid of full blocks
  const fullGridBlack = Math.floor((fullBlocksX * fullBlocksY) / 2);
  count += fullGridBlack * resolution * resolution;
  
  // Partial right column
  if (partialWidthX > 0) {
    const blackCountInColumn = fullBlocksX % 2 === 0 
      ? Math.floor(fullBlocksY / 2) 
      : Math.ceil(fullBlocksY / 2);
    count += blackCountInColumn * resolution * partialWidthX;
  }
  
  // Partial bottom row
  if (partialHeightY > 0) {
    const blackCountInRow = fullBlocksY % 2 === 0 
      ? Math.floor(fullBlocksX / 2) 
      : Math.ceil(fullBlocksX / 2);
    count += blackCountInRow * resolution * partialHeightY;
  }
  
  // Partial corner block
  if (partialWidthX > 0 && partialHeightY > 0) {
    if ((fullBlocksX + fullBlocksY) % 2 === 1) {
      count += partialWidthX * partialHeightY;
    }
  }
  
  return count;
}