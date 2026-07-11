function perimeter(n) {
  const fib = [1, 1];
  for (let i = 2; i <= n; i++) {
    fib.push(fib[i - 1] + fib[i - 2]);
  }
  return 4 * fib.slice(0, n + 1).reduce((sum, val) => sum + val, 0);
}