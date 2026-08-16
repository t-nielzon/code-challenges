function hasSubpattern(str) {
  const n = str.length;
  
  // Count character frequencies once
  const charCount = {};
  for (const char of str) {
    charCount[char] = (charCount[char] || 0) + 1;
  }
  
  // Try each possible subpattern length from 1 to n/2
  for (let patternLen = 1; patternLen <= n / 2; patternLen++) {
    // Pattern length must divide the string length evenly
    if (n % patternLen !== 0) continue;
    
    const repetitions = n / patternLen;
    
    // Check if all character counts are divisible by the number of repetitions
    let isValid = true;
    for (const count of Object.values(charCount)) {
      if (count % repetitions !== 0) {
        isValid = false;
        break;
      }
    }
    
    if (isValid) {
      return true;
    }
  }
  
  return false;
}