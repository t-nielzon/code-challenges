// solution.js
function firstNonRepeatingLetter(s) {
  const lowerStr = s.toLowerCase();
  const charCount = {};
  
  // count frequency of each character (case-insensitive)
  for (const char of lowerStr) {
    charCount[char] = (charCount[char] || 0) + 1;
  }
  
  // find first character with count of 1, preserving original case
  for (const char of s) {
    if (charCount[char.toLowerCase()] === 1) {
      return char;
    }
  }
  
  // no non-repeating character found
  return '';
}