function gap(g, m, n) {
  function isPrime(num) {
    if (num < 2) return false;
    if (num === 2) return true;
    if (num % 2 === 0) return false;
    for (let i = 3; i * i <= num; i += 2) {
      if (num % i === 0) return false;
    }
    return true;
  }
  
  let prevPrime = null;
  
  for (let i = m; i <= n; i++) {
    if (isPrime(i)) {
      if (prevPrime !== null && i - prevPrime === g) {
        return [prevPrime, i];
      }
      prevPrime = i;
    }
  }
  
  return null;
}