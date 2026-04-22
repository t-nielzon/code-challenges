function knapsack(capacity, items) {
  const indexed = items.map(([size, value], i) => ({ size, value, ratio: value / size, i }));
  indexed.sort((a, b) => b.ratio - a.ratio);
  const quantities = new Array(items.length).fill(0);
  let remaining = capacity;
  for (const { size, i } of indexed) {
    const count = Math.floor(remaining / size);
    if (count > 0) {
      quantities[i] = count;
      remaining -= count * size;
    }
  }
  return quantities;
}