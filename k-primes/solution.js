function countPrimeFactors(n) {
  if (n <= 1) return 0;
  
  let count = 0;
  
  while (n % 2 === 0) {
    count++;
    n = n / 2;
  }
  
  for (let i = 3; i * i <= n; i += 2) {
    while (n % i === 0) {
      count++;
      n = n / i;
    }
  }
  
  if (n > 1) {
    count++;
  }
  
  return count;
}

function countKprimes(k, start, end) {
  const result = [];
  
  for (let i = start; i <= end; i++) {
    if (countPrimeFactors(i) === k) {
      result.push(i);
    }
  }
  
  return result;
}

function puzzle(s) {
  const kprimes = {
    1: [],
    3: [],
    7: []
  };
  
  for (let i = 2; i <= s; i++) {
    const count = countPrimeFactors(i);
    if (count === 1) {
      kprimes[1].push(i);
    } else if (count === 3) {
      kprimes[3].push(i);
    } else if (count === 7) {
      kprimes[7].push(i);
    }
  }
  
  const sevenPrimes = new Set(kprimes[7]);
  
  let solutions = 0;
  
  for (let a of kprimes[1]) {
    for (let b of kprimes[3]) {
      const c = s - a - b;
      
      if (c > 0 && sevenPrimes.has(c)) {
        solutions++;
      }
    }
  }
  
  return solutions;
}