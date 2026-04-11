function pickPeaks(arr) {
  const result = { pos: [], peaks: [] };
  if (arr.length < 3) return result;

  let peakStart = null;

  for (let i = 1; i < arr.length - 1; i++) {
    if (arr[i] > arr[i - 1]) {
      peakStart = i;
    } else if (arr[i] < arr[i - 1]) {
      if (peakStart !== null) {
        result.pos.push(peakStart);
        result.peaks.push(arr[peakStart]);
        peakStart = null;
      }
    }
  }

  // handle plateau that ends at the last element — not a peak
  // (peakStart is intentionally discarded if never confirmed by a descent)

  return result;
}

module.exports = { pickPeaks };