function egyptianMultiplication(a, b) {
  let m = Math.max(a, b);
  let n = Math.min(a, b);
  
  const result = [];
  
  while (m > 0) {
    if (m % 2 === 1) {
      result.push(n);
    }
    m = Math.floor(m / 2);
    n *= 2;
  }
  
  return result.reverse();
}