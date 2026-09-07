function doublelyNotLess(n) {
  const nBig = BigInt(n);
  
  function reverseNumber(num) {
    const str = num.toString();
    const reversed = str.split('').reverse().join('');
    return BigInt(reversed);
  }
  
  let m = nBig;
  while (true) {
    const mReversed = reverseNumber(m);
    if (m >= nBig && mReversed >= nBig) {
      return m.toString();
    }
    m++;
  }
}