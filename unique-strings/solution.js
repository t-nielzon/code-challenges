function uniqCount(s) {
  const upper = s.toUpperCase();
  const freq = {};
  for (const ch of upper) {
    freq[ch] = (freq[ch] || 0) + 1;
  }

  const factorial = n => {
    let result = 1n;
    for (let i = 2n; i <= n; i++) result *= i;
    return result;
  };

  let result = factorial(BigInt(upper.length));
  for (const count of Object.values(freq)) {
    result /= factorial(BigInt(count));
  }
  return result;
}

module.exports = { uniqCount };