function prod2sum(a, b, c, d) {
  // brahmagupta–fibonacci identity: (a²+b²)(c²+d²) = (ac+bd)² + (ad-bc)² = (ac-bd)² + (ad+bc)²
  const pairs = new Set();
  const candidates = [
    [Math.abs(a * c + b * d), Math.abs(a * d - b * c)],
    [Math.abs(a * c - b * d), Math.abs(a * d + b * c)]
  ];
  for (const [e, f] of candidates) {
    const lo = Math.min(e, f);
    const hi = Math.max(e, f);
    pairs.add(`${lo},${hi}`);
  }
  return Array.from(pairs)
    .map(s => s.split(',').map(Number))
    .sort((a, b) => a[0] - b[0]);
}

module.exports = { prod2sum };