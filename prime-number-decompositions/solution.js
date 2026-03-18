function isValid(n) {
  return typeof n === 'number' && Number.isInteger(n) && n > 0;
}

function getAllPrimeFactors(n) {
  if (!isValid(n)) return [];
  if (n === 1) return [1];

  const factors = [];
  let d = 2;
  while (d * d <= n) {
    while (n % d === 0) {
      factors.push(d);
      n /= d;
    }
    d++;
  }
  if (n > 1) factors.push(n);
  return factors;
}

function getUniquePrimeFactorsWithCount(n) {
  if (!isValid(n)) return [[], []];
  if (n === 1) return [[1], [1]];

  const factors = getAllPrimeFactors(n);
  const primes = [];
  const counts = [];

  for (const f of factors) {
    if (primes.length && primes[primes.length - 1] === f) {
      counts[counts.length - 1]++;
    } else {
      primes.push(f);
      counts.push(1);
    }
  }
  return [primes, counts];
}

function getUniquePrimeFactorsWithProducts(n) {
  if (!isValid(n)) return [];
  if (n === 1) return [1];

  const [primes, counts] = getUniquePrimeFactorsWithCount(n);
  return primes.map((p, i) => Math.pow(p, counts[i]));
}