function minCoffeeSwaps(tray, deliverList) {
  // rotate the tray 90 degrees clockwise (works for non-square trays too)
  const rotate = (m) => {
    const rows = m.length, cols = m[0].length, res = [];
    for (let j = 0; j < cols; j++) {
      const row = [];
      for (let i = rows - 1; i >= 0; i--) row.push(m[i][j]);
      res.push(row);
    }
    return res;
  };

  // mirror horizontally; this also models picking coffees right-to-left
  const flip = (m) => m.map((r) => [...r].reverse());

  // read the tray row by row into a single sequence
  const flatten = (m) => m.reduce((acc, r) => acc.concat(r), []);

  // target index of every coffee in the deliverList
  const pos = new Map();
  deliverList.forEach((v, i) => pos.set(v, i));

  // minimum swaps to sort a sequence = n - (number of permutation cycles)
  const swapsFor = (arr) => {
    const n = arr.length;
    const target = arr.map((v) => pos.get(v));
    const seen = new Array(n).fill(false);
    let cycles = 0;
    for (let i = 0; i < n; i++) {
      if (seen[i]) continue;
      cycles++;
      let j = i;
      while (!seen[j]) {
        seen[j] = true;
        j = target[j];
      }
    }
    return n - cycles;
  };

  // generate the 8 dihedral orientations (4 rotations, each with/without mirror)
  const orientations = [];
  let cur = tray;
  for (let r = 0; r < 4; r++) {
    orientations.push(cur);
    orientations.push(flip(cur));
    cur = rotate(cur);
  }

  let best = Infinity;
  for (const o of orientations) {
    best = Math.min(best, swapsFor(flatten(o)));
  }
  return best;
}