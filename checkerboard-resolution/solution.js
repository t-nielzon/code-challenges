function countBlackSquares(width, height, resolution) {
  if (width === 0 || height === 0) return 0;
  
  const cols_full = Math.floor(width / resolution);
  const rows_full = Math.floor(height / resolution);
  const rem_w = width % resolution;
  const rem_h = height % resolution;
  
  let count = 0;
  
  // Interior full cells
  const blackCellsInInterior = countBlackCells(cols_full, rows_full);
  count += blackCellsInInterior * resolution * resolution;
  
  // Right strip (if exists)
  if (rem_w > 0) {
    const blackInRightStrip = cols_full % 2 === 0 ? Math.floor(rows_full / 2) : Math.ceil(rows_full / 2);
    count += blackInRightStrip * rem_w * resolution;
  }
  
  // Bottom strip (if exists)
  if (rem_h > 0) {
    const blackInBottomStrip = rows_full % 2 === 0 ? Math.floor(cols_full / 2) : Math.ceil(cols_full / 2);
    count += blackInBottomStrip * resolution * rem_h;
  }
  
  // Corner (if exists)
  if (rem_w > 0 && rem_h > 0) {
    if ((cols_full + rows_full) % 2 === 1) {
      count += rem_w * rem_h;
    }
  }
  
  return count;
}

function countBlackCells(cols, rows) {
  if (cols === 0 || rows === 0) return 0;
  const evenRows = Math.ceil(rows / 2);
  const oddRows = Math.floor(rows / 2);
  const oddCols = Math.floor(cols / 2);
  const evenCols = Math.ceil(cols / 2);
  return evenRows * oddCols + oddRows * evenCols;
}