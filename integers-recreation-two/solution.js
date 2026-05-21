function prod2sum(a, b, c, d) {
  // brahmagupta-fibonacci identity:
  // (a² + b²)(c² + d²) = (ac - bd)² + (ad + bc)² = (ac + bd)² + (ad - bc)²
  const pairs = [
    [Math.abs(a * c - b * d), Math.abs(a * d + b * c)],
    [Math.abs(a * c + b * d), Math.abs(a * d - b * c)]
  ].map(p => p.sort((x, y) => x - y));

  const seen = new Set();
  const result = [];
  for (const p of pairs) {
    const key = p.join(',');
    if (!seen.has(key)) {
      seen.add(key);
      result.push(p);
    }
  }
  return result.sort((x, y) => x[0] - y[0] || x[1] - y[1]);
}