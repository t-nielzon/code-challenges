function childrenAndApples(apples) {
  const n = apples.length;
  const total = apples.reduce((sum, a) => sum + a, 0);
  
  // Check if equal distribution is possible
  if (total % n !== 0) {
    return -1;
  }
  
  const target = total / n;
  
  // Calculate the difference from target for each child
  // Positive means excess, negative means deficit
  const differences = apples.map(a => a - target);
  
  // Check if all differences are divisible by 2
  // (since we can only transfer 2 apples at a time)
  for (const diff of differences) {
    if (diff % 2 !== 0) {
      return -1;
    }
  }
  
  // Count total excess (or deficit, they should be equal in magnitude)
  // Each step transfers 2 apples, reducing excess by 2 and deficit by 2
  // So total steps = sum of all positive differences / 2
  let totalExcess = 0;
  for (const diff of differences) {
    if (diff > 0) {
      totalExcess += diff;
    }
  }
  
  // Each step transfers 2 apples, so we need totalExcess / 2 steps
  return totalExcess / 2;
}