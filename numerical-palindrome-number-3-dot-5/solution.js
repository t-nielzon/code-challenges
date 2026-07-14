// solution.js
function numPalindrome(num) {
  if (!Number.isInteger(num) || num < 0) {
    return "Not valid";
  }
  
  const str = String(num);
  const palindromes = new Set();
  
  // Generate all substrings
  for (let i = 0; i < str.length; i++) {
    for (let j = i + 1; j <= str.length; j++) {
      const substring = str.substring(i, j);
      
      // Skip single digit numbers
      if (substring.length === 1) {
        continue;
      }
      
      // Skip if starts with zero
      if (substring[0] === '0') {
        continue;
      }
      
      // Check if palindrome
      const reversed = substring.split('').reverse().join('');
      if (substring === reversed) {
        palindromes.add(Number(substring));
      }
    }
  }
  
  if (palindromes.size === 0) {
    return "No palindromes found";
  }
  
  return Array.from(palindromes).sort((a, b) => a - b);
}