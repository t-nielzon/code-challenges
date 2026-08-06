function modpow(x, y, n) {
  let result = 1;
  x = x % n;
  
  while (y > 0) {
    if (y % 2 === 1) {
      result = (result * x) % n;
    }
    y = Math.floor(y / 2);
    x = (x * x) % n;
  }
  
  return result;
}