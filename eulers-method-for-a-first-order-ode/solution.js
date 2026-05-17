function ex_euler(n) {
  const h = 1 / n;
  const f = (x, y) => 2 - Math.exp(-4 * x) - 2 * y;
  const z = (x) => 1 + 0.5 * Math.exp(-4 * x) - 0.5 * Math.exp(-2 * x);

  let y = 1;
  let x = 0;
  let sum = 0;

  for (let k = 0; k <= n; k++) {
    const zk = z(x);
    const err = zk === 0 ? 0 : Math.abs(y - zk) / zk;
    sum += err;
    if (k < n) {
      y = y + f(x, y) * h;
      x = x + h;
    }
  }

  const mean = sum / (n + 1);
  return Math.trunc(mean * 1e6) / 1e6;
}