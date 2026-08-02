function consec_kprimes(k, arr) {
  function countPrimeFactors(n) {
    let count = 0;
    
    // Check for factor 2
    while (n % 2 === 0) {
      count++;
      n = n / 2;
    }
    
    // Check for odd factors from 3 onwards
    for (let i = 3; i * i <= n; i += 2) {
      while (n % i === 0) {
        count++;
        n = n / i;
      }
    }
    
    // If n is still > 1, it's a prime factor
    if (n > 1) {
      count++;
    }
    
    return count;
  }
  
  let count = 0;
  for (let i = 0; i < arr.length - 1; i++) {
    if (countPrimeFactors(arr[i]) === k && countPrimeFactors(arr[i + 1]) === k) {
      count++;
    }
  }
  return count;
}