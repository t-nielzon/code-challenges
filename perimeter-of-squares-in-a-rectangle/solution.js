function perimeter(n) {
  let a = 1n, b = 1n, sum = 1n;
  for (let i = 1; i <= n; i++) {
    sum += b;
    [a, b] = [b, a + b];
  }
  return 4n * sum;
}