function regexNumberLessThan(n) {
  const s = n.toString();
  const len = s.length;
  const parts = [];
  
  // Add patterns for all numbers with fewer digits than n
  for (let i = 1; i < len; i++) {
    if (i === 1) {
      parts.push('[1-9]');
    } else {
      parts.push('[1-9]' + '[0-9]'.repeat(i - 1));
    }
  }
  
  // Add patterns for numbers with same number of digits as n but strictly less than n
  for (let i = 0; i < s.length; i++) {
    const digit = parseInt(s[i]);
    
    if (digit === 0) {
      continue;
    }
    
    const prefix = s.substring(0, i);
    let range;
    
    if (i === 0) {
      // First digit must be between 1 and digit-1
      if (digit === 1) {
        continue;
      }
      range = '[1-' + (digit - 1) + ']';
    } else {
      // Other digits can be between 0 and digit-1
      if (digit === 1) {
        range = '0';
      } else {
        range = '[0-' + (digit - 1) + ']';
      }
    }
    
    const suffix = '[0-9]'.repeat(s.length - i - 1);
    parts.push(prefix + range + suffix);
  }
  
  if (parts.length === 0) {
    return '^$';
  }
  
  return '^(' + parts.join('|') + ')$';
}