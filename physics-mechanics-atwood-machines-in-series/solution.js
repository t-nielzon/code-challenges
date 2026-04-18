function atwood(masses) {
  const g = 9.80665;
  const reduce = (arr) => {
    if (arr.length === 2) return arr[1];
    const [m2, ...rest] = arr;
    const m3 = reduce(rest);
    return (4 * m2 * m3) / (m2 + m3);
  };
  const m1 = masses[0];
  const m2eff = reduce(masses);
  return g * (m2eff - m1) / (m1 + m2eff);
}