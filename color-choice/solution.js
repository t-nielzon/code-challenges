function chooseColour(m, n) {
  if (m === 1) return 0;
  let c = 1n;
  const N = BigInt(n);
  const M = BigInt(m);
  for (let x = 1n; x <= N; x++) {
    c = c * (N - x + 1n) / x;
    if (c === M) return Number(x);
    if (c > M) return -1;
  }
  return -1;
}