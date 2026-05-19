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

function consecKprimes(k, arr) {
  let result = 0;
  for (let i = 1; i < arr.length; i++) {
    if (countPrimeFactors(arr[i - 1]) === k && countPrimeFactors(arr[i]) === k) {
      result++;
    }
  }
  return result;
}