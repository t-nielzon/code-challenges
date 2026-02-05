function solve(s) {
  let count = 0;
  
  // a substring ending at position i is odd if and only if the digit at position i is odd
  // for each odd digit at position i, there are (i + 1) substrings ending at that position
  for (let i = 0; i < s.length; i++) {
    if (parseInt(s[i]) % 2 !== 0) {
      count += i + 1;
    }
  }
  
  return count;
}