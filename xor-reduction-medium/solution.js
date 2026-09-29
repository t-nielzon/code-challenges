function xorReduction(m, n) {
  function xorFrom0ToN(num) {
    if (num < 0) return 0;
    switch (num % 4) {
      case 0: return num;
      case 1: return 1;
      case 2: return num + 1;
      case 3: return 0;
    }
  }
  
  return xorFrom0ToN(n) ^ xorFrom0ToN(m - 1);
}