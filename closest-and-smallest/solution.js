function closest(strng) {
  if (strng === "") return [];

  const weight = (n) =>
    String(n)
      .split("")
      .reduce((sum, d) => sum + Number(d), 0);

  // build [weight, index, originalNumber] for each entry, then sort by
  // weight then index so neighbouring pairs are candidates for "closest".
  const items = strng
    .split(" ")
    .map((num, idx) => [weight(num), idx, Number(num)])
    .sort((a, b) => a[0] - b[0] || a[1] - b[1]);

  let best = null;
  for (let i = 1; i < items.length; i++) {
    const diff = items[i][0] - items[i - 1][0];
    if (best === null || diff < best.diff) {
      best = { diff, pair: [items[i - 1], items[i]] };
    }
  }

  return best.pair;
}