function simpson(n) {
  const a = 0;
  const b = Math.PI;
  const h = (b - a) / n;
  const f = (x) => 1.5 * Math.pow(Math.sin(x), 3);

  let sumOdd = 0;
  for (let i = 1; i <= n / 2; i++) {
    sumOdd += f(a + (2 * i - 1) * h);
  }

  let sumEven = 0;
  for (let i = 1; i <= n / 2 - 1; i++) {
    sumEven += f(a + 2 * i * h);
  }

  return ((b - a) / (3 * n)) * (f(a) + f(b) + 4 * sumOdd + 2 * sumEven);
}