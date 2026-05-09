function knapsack(capacity, items) {
  const indexed = items.map(([size, value], i) => ({ size, value, ratio: value / size, i }));
  indexed.sort((a, b) => b.ratio - a.ratio);
  const quantities = new Array(items.length).fill(0);
  let remaining = capacity;
  for (const { size, i } of indexed) {
    const qty = Math.floor(remaining / size);
    quantities[i] = qty;
    remaining -= qty * size;
  }
  return quantities;
}