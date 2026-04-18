function interp(f, l, u, n) {
  const d = (u - l) / n;
  const result = [];
  for (let i = 0; i < n; i++) {
    const y = f(l + i * d);
    result.push(Math.floor(y * 100.0) / 100.0);
  }
  return result;
}