function primeFactorCount(n) {
  let count = 0;
  let d = 2;
  while (d * d <= n) {
    while (n % d === 0) {
      n /= d;
      count++;
    }
    d++;
  }
  if (n > 1) count++;
  return count;
}

function countKprimes(k, start, nd) {
  const result = [];
  for (let n = start; n <= nd; n++) {
    if (primeFactorCount(n) === k) result.push(n);
  }
  return result;
}

function puzzle(s) {
  // a is 1-prime, b is 3-prime, c is 7-prime, a + b + c = s
  const ones = countKprimes(1, 2, s);
  const threes = new Set(countKprimes(3, 1, s));
  const sevens = new Set(countKprimes(7, 1, s));

  let total = 0;
  for (const a of ones) {
    for (const c of sevens) {
      const b = s - a - c;
      if (b > 0 && threes.has(b)) total++;
    }
  }
  return total;
}