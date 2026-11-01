function primeFactors(n) {
  let result = '';
  let divisor = 2;
  
  while (divisor * divisor <= n) {
    let count = 0;
    while (n % divisor === 0) {
      count++;
      n /= divisor;
    }
    if (count > 0) {
      result += '(' + divisor;
      if (count > 1) {
        result += '**' + count;
      }
      result += ')';
    }
    divisor++;
  }
  
  if (n > 1) {
    result += '(' + n + ')';
  }
  
  return result;
}