function* paperFold() {
  const memo = {};
  
  function getValue(n) {
    if (n in memo) return memo[n];
    
    let result;
    if (n % 2 === 0) {
      // even positions refer to earlier positions: a(2n) = a(n)
      result = getValue(n / 2);
    } else if (n % 4 === 1) {
      // odd positions where n ≡ 1 (mod 4) are always 1
      result = 1;
    } else {
      // odd positions where n ≡ 3 (mod 4) are always 0
      result = 0;
    }
    memo[n] = result;
    return result;
  }
  
  let n = 1;
  while (true) {
    yield getValue(n);
    n++;
  }
}