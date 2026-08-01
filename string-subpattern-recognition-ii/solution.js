function check(s) {
  const charCount = {};
  for (const char of s) {
    charCount[char] = (charCount[char] || 0) + 1;
  }
  
  // GCD of all character counts equals the number of repetitions needed
  const gcd = (a, b) => b === 0 ? a : gcd(b, a % b);
  
  let g = 0;
  for (const count of Object.values(charCount)) {
    g = gcd(g, count);
  }
  
  return g > 1;
}