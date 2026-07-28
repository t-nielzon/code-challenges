function decomp(n) {
  function sieve(max) {
    if (max < 2) return [];
    const isPrime = Array(max + 1).fill(true);
    isPrime[0] = isPrime[1] = false;
    
    for (let i = 2; i * i <= max; i++) {
      if (isPrime[i]) {
        for (let j = i * i; j <= max; j += i) {
          isPrime[j] = false;
        }
      }
    }
    
    const primes = [];
    for (let i = 2; i <= max; i++) {
      if (isPrime[i]) primes.push(i);
    }
    return primes;
  }
  
  function countPrimesInFactorial(n, p) {
    let count = 0;
    let num = n;
    while (num >= p) {
      num = Math.floor(num / p);
      count += num;
    }
    return count;
  }
  
  const primes = sieve(n);
  const result = [];
  
  for (const prime of primes) {
    const exponent = countPrimesInFactorial(n, prime);
    if (exponent > 0) {
      if (exponent === 1) {
        result.push(`${prime}`);
      } else {
        result.push(`${prime}^${exponent}`);
      }
    }
  }
  
  return result.join(' * ');
}