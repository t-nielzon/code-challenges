function countPrimeFactors(n) {
  let count = 0;
  for (let p = 2; p * p <= n; p++) {
    while (n % p === 0) {
      count++;
      n = n / p;
    }
  }
  if (n > 1) count++;
  return count;
}

function countKprimes(k, start, nd) {
  const result = [];
  for (let n = start; n <= nd; n++) {
    if (countPrimeFactors(n) === k) result.push(n);
  }
  return result;
}

function puzzle(s) {
  const ones = countKprimes(1, 2, s);
  const threes = countKprimes(3, 8, s);
  const sevens = countKprimes(7, 128, s);
  const sevenSet = new Set(sevens);
  let total = 0;
  for (const a of ones) {
    if (a >= s) break;
    for (const b of threes) {
      const c = s - a - b;
      if (c < 128) break;
      if (sevenSet.has(c)) total++;
    }
  }
  return total;
}