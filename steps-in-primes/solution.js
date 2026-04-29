function step(g, m, n) {
  const isPrime = (k) => {
    if (k < 2) return false;
    if (k % 2 === 0) return k === 2;
    if (k % 3 === 0) return k === 3;
    for (let i = 5; i * i <= k; i += 6) {
      if (k % i === 0 || k % (i + 2) === 0) return false;
    }
    return true;
  };
  for (let i = m; i <= n - g; i++) {
    if (isPrime(i) && isPrime(i + g)) return [i, i + g];
  }
  return null;
}