function gap(g, m, n) {
  const isPrime = (x) => {
    if (x < 2) return false;
    if (x % 2 === 0) return x === 2;
    for (let i = 3; i * i <= x; i += 2) {
      if (x % i === 0) return false;
    }
    return true;
  };

  let prev = null;
  for (let i = m; i <= n; i++) {
    if (isPrime(i)) {
      if (prev !== null && i - prev === g) {
        return [prev, i];
      }
      prev = i;
    }
  }
  return null;
}