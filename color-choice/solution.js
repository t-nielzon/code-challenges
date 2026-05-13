function colorChoice(m, n) {
  const M = BigInt(m);
  const N = BigInt(n);
  if (M === 1n) return 0;
  let c = 1n;
  for (let x = 1n; x <= N; x++) {
    c = c * (N - x + 1n) / x;
    if (c === M) return Number(x);
    if (c > M) return -1;
  }
  return -1;
}