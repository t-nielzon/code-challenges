function buddy(start, limit) {
  function sumOfDivisors(n) {
    let sum = 1;
    for (let i = 2; i * i <= n; i++) {
      if (n % i === 0) {
        sum += i;
        if (i !== n / i) {
          sum += n / i;
        }
      }
    }
    return sum;
  }
  
  for (let n = start; n <= limit; n++) {
    const sn = sumOfDivisors(n);
    const m = sn - 1;
    
    if (m > n && sumOfDivisors(m) === n + 1) {
      return [n, m];
    }
  }
  
  return "Nothing";
}