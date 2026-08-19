function decomp(n) {
  function sieve(limit) {
    if (limit < 2) return [];
    const isPrime = new Array(limit + 1).fill(true);
    isPrime[0] = isPrime[1] = false;
    
    for (let i = 2; i * i <= limit; i++) {
      if (isPrime[i]) {
        for (let j = i * i; j <= limit; j += i) {
          isPrime[j] = false;
        }
      }
    }
    
    const primes = [];
    for (let i = 2; i <= limit; i++) {
      if (isPrime[i]) primes.push(i);
    }
    return primes;
  }
  
  function countPrimeInFactorial(n, p) {
    let count = 0;
    let power = p;
    while (power <= n) {
      count += Math.floor(n / power);
      power *= p;
    }
    return count;
  }
  
  const primes = sieve(n);
  const factors = [];
  
  for (const prime of primes) {
    const count = countPrimeInFactorial(n, prime);
    if (count > 0) {
      if (count === 1) {
        factors.push(prime.toString());
      } else {
        factors.push(`${prime}^${count}`);
      }
    }
  }
  
  return factors.join(' * ');
}