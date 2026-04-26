function ex_euler(n) {
  const h = 1 / n;
  const f = (x, y) => 2 - Math.exp(-4 * x) - 2 * y;
  const exact = (x) => 1 + 0.5 * Math.exp(-4 * x) - 0.5 * Math.exp(-2 * x);

  let y = 1;
  let sum = 0;
  for (let k = 0; k <= n; k++) {
    const x = k * h;
    const z = exact(x);
    sum += Math.abs(y - z) / z;
    y = y + f(x, y) * h;
  }
  const mean = sum / (n + 1);
  return Math.trunc(mean * 1e6) / 1e6;
}