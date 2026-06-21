function perimeter(n) {
  let a = 1n;
  let b = 1n;
  let sum = 1n;
  for (let i = 1; i <= n; i++) {
    sum += b;
    [a, b] = [b, a + b];
  }
  return 4n * sum;
}