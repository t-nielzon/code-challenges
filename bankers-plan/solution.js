function fortune(f0, p, c0, n, i) {
  let f = f0;
  let c = c0;
  const pRate = p / 100;
  const iRate = i / 100;

  for (let year = 2; year <= n; year++) {
    f = Math.trunc(f + f * pRate - c);
    c = Math.trunc(c + c * iRate);
    if (f < 0) return false;
  }
  return true;
}

module.exports = { fortune };