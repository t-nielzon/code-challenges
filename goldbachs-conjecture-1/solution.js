function goldbach(n) {
  if (n % 2 !== 0) return [];

  function isPrime(num) {
    if (num < 2) return false;
    if (num < 4) return true;
    if (num % 2 === 0 || num % 3 === 0) return false;
    for (let i = 5; i * i <= num; i += 6) {
      if (num % i === 0 || num % (i + 2) === 0) return false;
    }
    return true;
  }

  const result = [];
  for (let i = 2; i <= n / 2; i++) {
    if (isPrime(i) && isPrime(n - i)) {
      result.push(`${i}+${n - i}`);
    }
  }
  return result;
}

module.exports = { goldbach };