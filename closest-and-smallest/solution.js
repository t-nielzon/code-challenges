function closest(strng) {
  if (strng === "") return [];
  const weight = (s) => s.split("").reduce((a, c) => a + Number(c), 0);
  const items = strng.split(" ").map((n, i) => [weight(n), i, n]);
  const sorted = items.slice().sort((a, b) => a[0] - b[0] || a[1] - b[1]);
  let best = null;
  for (let i = 1; i < sorted.length; i++) {
    const diff = sorted[i][0] - sorted[i - 1][0];
    if (best === null || diff < best.diff) {
      best = { diff, pair: [sorted[i - 1], sorted[i]] };
    }
  }
  return [
    [best.pair[0][0], best.pair[0][1], Number(best.pair[0][2])],
    [best.pair[1][0], best.pair[1][1], Number(best.pair[1][2])],
  ];
}