function cycleLength(n) {
  function gcd(a, b) {
    while (b !== 0) {
      const temp = b;
      b = a % b;
      a = temp;
    }
    return a;
  }
  
  if (gcd(n, 10) !== 1) {
    return -1;
  }
  
  const remainders = new Map();
  let remainder = 1;
  let position = 0;
  
  while (!remainders.has(remainder)) {
    remainders.set(remainder, position);
    remainder = (remainder * 10) % n;
    position++;
  }
  
  return position - remainders.get(remainder);
}