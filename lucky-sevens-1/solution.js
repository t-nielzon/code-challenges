function luckySevens(arr) {
  let count = 0;
  const rows = arr.length;
  const cols = arr[0].length;

  for (let r = 0; r < rows; r++) {
    for (let c = 0; c < cols; c++) {
      if (arr[r][c] !== 7) continue;

      const neighbors = [
        r > 0 ? arr[r - 1][c] : 0,
        r < rows - 1 ? arr[r + 1][c] : 0,
        c > 0 ? arr[r][c - 1] : 0,
        c < cols - 1 ? arr[r][c + 1] : 0,
      ];

      const sum = neighbors.reduce((s, v) => s + (typeof v === 'number' ? v : 0), 0);
      const cbrt = Math.round(Math.cbrt(sum));
      if (cbrt * cbrt * cbrt === sum) count++;
    }
  }

  return count;
}

module.exports = { luckySevens };