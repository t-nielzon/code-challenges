function findTrue(matrix) {
  // binary search to find the true element in an infinite proxy-based matrix
  // we can't iterate, so we search by narrowing ranges exponentially

  function findRow(mat) {
    // first find an upper bound where the row contains true
    let hi = 1;
    while (!rowHasTrue(mat, hi)) hi *= 2;

    // binary search for the exact row
    let lo = 0;
    while (lo < hi) {
      const mid = (lo + hi) >>> 1;
      if (rangeHasTrue(mat, 0, mid)) {
        hi = mid;
      } else {
        lo = mid + 1;
      }
    }
    return lo;
  }

  // check if any row from 0..maxRow has true in any column up to some bound
  function rowHasTrue(mat, maxRow) {
    for (let r = 0; r <= maxRow; r++) {
      if (mat[r][0] === true) return true;
    }
    // the true column might not be 0, so we need a different approach
    return rangeHasTrue(mat, 0, maxRow);
  }

  function rangeHasTrue(mat, lo, hi) {
    // check if any row in [lo, hi] contains true by searching columns
    for (let r = lo; r <= hi; r++) {
      if (searchCol(mat[r])) return true;
    }
    return false;
  }

  function searchCol(row) {
    // find upper bound for column
    let hi = 1;
    while (hi < 1e15) {
      if (colRangeHasTrue(row, 0, hi)) return true;
      hi *= 2;
    }
    return false;
  }

  function colRangeHasTrue(row, lo, hi) {
    for (let c = lo; c <= hi; c++) {
      if (row[c] === true) return true;
    }
    return false;
  }

  // better approach: binary search on rows, then binary search on columns
  // to find the row, probe row i for "has any true" by scanning columns
  // but scanning columns is also infinite...

  // key insight: use diagonal search. Check mat[i][j] along expanding squares
  // or use the fact that we can test individual cells directly.

  // most efficient: search expanding shells (0,0), (0,1),(1,0),(1,1), etc.
  // but that's O(n^2). Instead, find row first, then column, each via binary search.

  // to check if a row is the "true row", we need to find any true in it.
  // since the row is infinite, we binary search the column within it.
  // but we don't know if the row has true at all — we'd search forever.

  // approach: expand search area as a square, doubling size
  let size = 1;
  let y, x;

  outer:
  while (true) {
    for (let r = 0; r < size; r++) {
      for (let c = 0; c < size; c++) {
        if (matrix[r][c] === true) {
          return [r, c];
        }
      }
    }
    size *= 2;
  }
}