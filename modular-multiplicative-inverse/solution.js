function inverseMod(a, n) {
  function extendedGCD(a, b) {
    if (b === 0) {
      return [a, 1, 0];
    }
    const [gcd, x1, y1] = extendedGCD(b, a % b);
    const x = y1;
    const y = x1 - Math.floor(a / b) * y1;
    return [gcd, x, y];
  }
  
  const [gcd, x] = extendedGCD(a, n);
  
  if (gcd !== 1) {
    return null;
  }
  
  return (x % n + n) % n;
}