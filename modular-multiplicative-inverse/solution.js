function inverseMod(a, n) {
  function extGcd(a, b) {
    if (b === 0) {
      return { gcd: a, x: 1, y: 0 };
    }
    const { gcd, x: x1, y: y1 } = extGcd(b, a % b);
    const x = y1;
    const y = x1 - Math.floor(a / b) * y1;
    return { gcd, x, y };
  }
  
  const { gcd, x } = extGcd(a, n);
  
  if (gcd !== 1) {
    return null;
  }
  
  return (x % n + n) % n;
}