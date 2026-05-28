function closest(strng) {
  if (strng === "") return [];
  const items = strng.split(" ").map((n, i) => {
    const weight = n.split("").reduce((s, d) => s + Number(d), 0);
    return [weight, i, n];
  });
  const sorted = items.slice().sort((a, b) => a[0] - b[0] || a[1] - b[1]);
  let best = null;
  for (let i = 1; i < sorted.length; i++) {
    const diff = sorted[i][0] - sorted[i - 1][0];
    if (best === null || diff < best.diff) {
      best = { diff, pair: [sorted[i - 1], sorted[i]] };
    }
  }
  return best.pair;
}