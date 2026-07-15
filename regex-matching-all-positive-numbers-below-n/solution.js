function regex(n) {
  const nStr = String(n);
  const len = nStr.length;
  const patterns = [];
  
  // Patterns for all numbers with fewer digits than n
  for (let l = 1; l < len; l++) {
    if (l === 1) {
      patterns.push('[1-9]');
    } else {
      patterns.push('[1-9][0-9]{' + (l - 1) + '}');
    }
  }
  
  // Patterns for numbers with the same number of digits as n, but less than n
  const sameDigitPatterns = buildSameDigit(nStr);
  if (sameDigitPatterns.length > 0) {
    if (sameDigitPatterns.length === 1) {
      patterns.push(sameDigitPatterns[0]);
    } else {
      patterns.push('(?:' + sameDigitPatterns.join('|') + ')');
    }
  }
  
  let result;
  if (patterns.length === 0) {
    result = '(?!)';
  } else if (patterns.length === 1) {
    result = patterns[0];
  } else {
    result = '(?:' + patterns.join('|') + ')';
  }
  
  return '^' + result + '$';
}

function buildSameDigit(nStr) {
  const result = [];
  const len = nStr.length;
  
  function build(pos, current) {
    if (pos === len) {
      if (current !== nStr) {
        result.push(current);
      }
      return;
    }
    
    const digit = parseInt(nStr[pos]);
    
    // Match a digit less than current, followed by any digits
    if (digit > 0) {
      const from = pos === 0 ? 1 : 0;
      let digitPat;
      if (from === digit - 1) {
        digitPat = String(digit - 1);
      } else {
        digitPat = '[' + from + '-' + (digit - 1) + ']';
      }
      
      if (pos === len - 1) {
        result.push(current + digitPat);
      } else {
        result.push(current + digitPat + '[0-9]{' + (len - pos - 1) + '}');
      }
    }
    
    // Match current digit and continue
    build(pos + 1, current + digit);
  }
  
  build(0, '');
  
  return result;
}