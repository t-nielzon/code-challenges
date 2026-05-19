function step(g, m, n) {
  const isPrime = (num) => {
    if (num < 2) return false;
    if (num < 4) return true;
    if (num % 2 === 0) return false;
    for (let i = 3, lim = Math.sqrt(num); i <= lim; i += 2) {
      if (num % i === 0) return false;
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