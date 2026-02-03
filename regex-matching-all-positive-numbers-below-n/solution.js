function regexRangeBelow(n) {
  if (n <= 1) {
    // No positive numbers less than 1
    return "^$a"; // Pattern that matches nothing
  }
  
  if (n === 2) {
    return "^1$";
  }
  
  const nStr = String(n);
  const len = nStr.length;
  const patterns = [];
  
  // Pattern for numbers with fewer digits than n (1 to len-1 digits)
  if (len > 1) {
    if (len === 2) {
      patterns.push("[1-9]");
    } else {
      patterns.push("[1-9][0-9]{0," + (len - 2) + "}");
    }
  }
  
  // Pattern for numbers with same number of digits as n but less than n
  const sameDigitPatterns = buildSameDigitPatterns(nStr);
  patterns.push(...sameDigitPatterns);
  
  if (patterns.length === 0) {
    return "^$a"; // Match nothing
  }
  
  return "^(" + patterns.join("|") + ")$";
}

function buildSameDigitPatterns(nStr) {
  const patterns = [];
  const len = nStr.length;
  
  // For each position, create pattern where that position is less than n's digit
  // and all previous positions match n's digits exactly
  for (let i = 0; i < len; i++) {
    const digit = parseInt(nStr[i]);
    
    // Determine the range of valid digits at position i
    let minDigit, maxDigit;
    
    if (i === 0) {
      // First digit: must be >= 1 and < digit
      minDigit = 1;
      maxDigit = digit - 1;
    } else {
      // Other digits: must be >= 0 and < digit
      minDigit = 0;
      maxDigit = digit - 1;
    }
    
    if (maxDigit >= minDigit) {
      let pattern = "";
      
      // Add prefix (exact match of previous digits)
      for (let j = 0; j < i; j++) {
        pattern += nStr[j];
      }
      
      // Add the range for current position
      pattern += charRange(minDigit, maxDigit);
      
      // Add suffix (any digits for remaining positions)
      const remaining = len - i - 1;
      if (remaining > 0) {
        if (remaining === 1) {
          pattern += "[0-9]";
        } else {
          pattern += "[0-9]{" + remaining + "}";
        }
      }
      
      patterns.push(pattern);
    }
  }
  
  return patterns;
}

function charRange(min, max) {
  if (min > max) return "";
  if (min === max) return String(min);
  return "[" + min + "-" + max + "]";
}

module.exports = { regexRangeBelow };