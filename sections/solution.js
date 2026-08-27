function c(k) {
  // Check if k is a perfect square
  let m = Math.floor(Math.sqrt(k));
  if (m * m === k) {
    // m is correct
  } else if ((m + 1) * (m + 1) === k) {
    m = m + 1;
  } else {
    return 0;
  }
  
  // Prime factorize m
  const factors = {};
  let n = m;
  
  for (let i = 2; i * i <= n; i++) {
    while (n % i === 0) {
      factors[i] = (factors[i] || 0) + 1;
      n /= i;
    }
  }
  if (n > 1) {
    factors[n] = (factors[n] || 0) + 1;
  }
  
  // Count divisors of m^3
  let count = 1;
  for (const exp of Object.values(factors)) {
    count *= (3 * exp + 1);
  }
  
  return count;
}