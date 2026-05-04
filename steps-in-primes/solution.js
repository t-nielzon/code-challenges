function step(g, m, n) {
  const isPrime = (k) => {
    if (k < 2) return false;
    if (k < 4) return true;
    if (k % 2 === 0) return false;
    for (let i = 3; i * i <= k; i += 2) {
      if (k % i === 0) return false;
    }
    return true;
  };

  for (let i = m; i <= n - g; i++) {
    if (isPrime(i) && isPrime(i + g)) {
      return [i, i + g];
    }
  }
  return null;
}