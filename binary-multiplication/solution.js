function egyptianMultiplication(m, n) {
  if (m < n) {
    [m, n] = [n, m];
  }
  
  const list = [];
  
  while (m > 0) {
    if (m % 2 !== 0) {
      list.push(n);
    }
    m = Math.floor(m / 2);
    n = n * 2;
  }
  
  return list.reverse();
}