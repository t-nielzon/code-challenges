function modpow(x, y, n) {
  if (n === 1) return 0;
  if (y === 0) return 1;
  
  let result = 1;
  x = x % n;
  
  while (y > 0) {
    if (y & 1) {
      result = (result * x) % n;
    }
    y = y >> 1;
    x = (x * x) % n;
  }
  
  return result;
}