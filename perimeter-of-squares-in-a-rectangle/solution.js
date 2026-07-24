function perimeter(n) {
  let fib = [1, 1];
  for (let i = 2; i <= n; i++) {
    fib[i] = fib[i - 1] + fib[i - 2];
  }
  
  let sum = fib.slice(0, n + 1).reduce((a, b) => a + b, 0);
  return 4 * sum;
}