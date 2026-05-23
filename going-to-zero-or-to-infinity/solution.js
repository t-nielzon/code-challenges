function going(n) {
  let sum = 1;
  let term = 1;
  for (let k = n; k > 1; k--) {
    term /= k;
    sum += term;
  }
  return Math.floor(sum * 1e6) / 1e6;
}