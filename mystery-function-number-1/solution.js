function solved(x) {
  const n = x.length;
  const half = Math.ceil(n / 2);
  let result = '';
  for (let i = 0; i < half; i++) {
    result += x[i];
    if (i + half < n) result += x[i + half];
  }
  return result;
}