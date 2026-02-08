function uniqCount(s) {
  const lower = s.toLowerCase();
  const freq = {};
  for (const ch of lower) {
    freq[ch] = (freq[ch] || 0) + 1;
  }

  const factorial = (n) => {
    let result = 1n;
    for (let i = 2n; i <= n; i++) {
      result *= i;
    }
    return result;
  };

  let result = factorial(BigInt(lower.length));
  for (const count of Object.values(freq)) {
    result /= factorial(BigInt(count));
  }

  return result;
}

module.exports = { uniqCount };