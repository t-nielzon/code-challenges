function doubleNotLess(n) {
  let m = BigInt(n);
  const nBig = BigInt(n);
  
  while (true) {
    const mStr = m.toString();
    const reversed = BigInt(mStr.split('').reverse().join(''));
    
    if (reversed >= nBig) {
      return m.toString();
    }
    m += 1n;
  }
}