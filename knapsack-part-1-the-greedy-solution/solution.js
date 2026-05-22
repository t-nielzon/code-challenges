function knapsack(capacity, items) {
  const quantities = items.map(() => 0);
  let remaining = capacity;

  while (true) {
    let bestIndex = -1;
    let bestRatio = 0;

    for (let i = 0; i < items.length; i++) {
      const [size, value] = items[i];
      if (size <= remaining) {
        const ratio = value / size;
        if (ratio > bestRatio) {
          bestRatio = ratio;
          bestIndex = i;
        }
      }
    }

    if (bestIndex === -1) break;

    quantities[bestIndex]++;
    remaining -= items[bestIndex][0];
  }

  return quantities;
}