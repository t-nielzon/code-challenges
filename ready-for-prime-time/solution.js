function prime(num) {
  if (num < 2) return [];
  const sieve = new Array(num + 1).fill(true);
  sieve[0] = sieve[1] = false;
  for (let i = 2; i * i <= num; i++) {
    if (sieve[i]) {
      for (let j = i * i; j <= num; j += i) {
        sieve[j] = false;
      }
    }
  }
  const primes = [];
  for (let i = 2; i <= num; i++) {
    if (sieve[i]) primes.push(i);
  }
  return primes;
}

module.exports = prime;