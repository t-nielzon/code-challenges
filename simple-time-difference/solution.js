function maxInterval(arr) {
  const minutes = [...new Set(arr)].map(t => {
    const [h, m] = t.split(':').map(Number);
    return h * 60 + m;
  }).sort((a, b) => a - b);

  let maxGap = 0;

  for (let i = 1; i < minutes.length; i++) {
    maxGap = Math.max(maxGap, minutes[i] - minutes[i - 1] - 1);
  }

  // wrap-around gap
  maxGap = Math.max(maxGap, (24 * 60 - minutes[minutes.length - 1]) + minutes[0] - 1);

  const h = Math.floor(maxGap / 60);
  const m = maxGap % 60;
  return `${String(h).padStart(2, '0')}:${String(m).padStart(2, '0')}`;
}