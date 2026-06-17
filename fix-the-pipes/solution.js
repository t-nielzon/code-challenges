function fixThePipes(map, start, end) {
  const rows = map.length;
  const cols = map[0].length;
  const grid = map.map(r => r.split(''));

  // each column holds one contiguous vertical run of 'x'; record its endpoints
  const tops = [];
  const bottoms = [];
  for (let c = 0; c < cols; c++) {
    let t = -1, b = -1;
    for (let r = 0; r < rows; r++) {
      if (grid[r][c] === 'x') {
        if (t === -1) t = r;
        b = r;
      }
    }
    tops[c] = t;
    bottoms[c] = b;
  }

  // pick the box-drawing glyph from the set of connected directions
  function pipeChar(L, R, U, D) {
    if (L && R) return '━';
    if (U && D) return '┃';
    if (U && R) return '┗';
    if (D && L) return '┓';
    if (D && R) return '┏';
    if (U && L) return '┛';
  }

  // the path is monotonic across columns: it enters a column at one endpoint
  // of that column's vertical run and leaves at the other, then connects
  // horizontally into the next column at the same row
  let entry = start;
  for (let c = 0; c < cols; c++) {
    const t = tops[c], b = bottoms[c];
    const exit = entry === t ? b : t;

    for (let r = t; r <= b; r++) {
      let L = false, R = false, U = false, D = false;

      if (t === b) {
        // single cell just passes straight through
        L = true;
        R = true;
      } else {
        if (r > t) U = true;
        if (r < b) D = true;
        if (r === entry) L = true; // connects left (previous column or source)
        if (r === exit) R = true;  // connects right (next column or end)
      }

      grid[r][c] = pipeChar(L, R, U, D);
    }

    entry = exit; // next column enters where this one exited
  }

  return grid.map(r => r.join(''));
}