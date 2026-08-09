function mirroredExponentialChunks(arr) {
  if (arr.length === 0) return [];
  if (arr.length === 1) return [arr];
  
  const result = [];
  let idx = 0;
  let remaining = arr.length;
  const levels = [];
  
  const hasMedian = remaining % 2 === 1;
  if (hasMedian) {
    levels.push(1);
    remaining -= 1;
  }
  
  let levelExponent = 1;
  while (remaining > 0) {
    const size = Math.pow(2, levelExponent);
    if (remaining >= 2 * size) {
      levels.push(size);
      levels.push(size);
      remaining -= 2 * size;
    } else if (remaining > 0) {
      const size_partial = remaining / 2;
      levels.push(size_partial);
      levels.push(size_partial);
      remaining = 0;
    }
    levelExponent++;
  }
  
  const numLevels = (levels.length - (hasMedian ? 1 : 0)) / 2;
  
  for (let i = numLevels; i >= 1; i--) {
    const size = hasMedian ? levels[2 * i - 1] : levels[2 * i - 2];
    result.push(arr.slice(idx, idx + size));
    idx += size;
  }
  
  if (hasMedian) {
    result.push(arr.slice(idx, idx + 1));
    idx += 1;
  }
  
  for (let i = 1; i <= numLevels; i++) {
    const size = hasMedian ? levels[2 * i] : levels[2 * i - 1];
    result.push(arr.slice(idx, idx + size));
    idx += size;
  }
  
  return result;
}