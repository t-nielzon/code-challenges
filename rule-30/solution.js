function rule30(cells, n) {
  // rule 30 lookup: index is the 3-bit pattern (left, center, right), value is the new cell
  const rule = [0, 1, 1, 1, 1, 0, 0, 0];

  let current = cells;

  for (let i = 0; i < n; i++) {
    // pad with zeros on both sides to allow growth
    const padded = [0, 0, ...current, 0, 0];
    const next = [];

    for (let j = 1; j < padded.length - 1; j++) {
      const pattern = (padded[j - 1] << 2) | (padded[j] << 1) | padded[j + 1];
      next.push(rule[pattern]);
    }

    current = next;
  }

  return current;
}

module.exports = { rule30 };