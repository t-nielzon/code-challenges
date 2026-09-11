function xorReduction(m, n) {
  // compute xor from 0 to x using pattern: xor(0..x) depends on x % 4
  function xorUpTo(x) {
    if (x < 0) return 0;
    const remainder = x % 4;
    if (remainder === 0) return x;
    if (remainder === 1) return 1;
    if (remainder === 2) return x + 1;
    return 0;
  }
  
  // xor(m..n) = xor(0..n) ^ xor(0..m-1)
  return xorUpTo(n) ^ xorUpTo(m - 1);
}