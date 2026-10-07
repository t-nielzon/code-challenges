function unsquareDigits(num) {
  const str = String(num);
  const results = [];
  const squares = {0: 0, 1: 1, 4: 2, 9: 3, 16: 4, 25: 5, 36: 6, 49: 7, 64: 8, 81: 9};
  
  function backtrack(index, current) {
    if (index === str.length) {
      results.push(Number(current));
      return;
    }
    
    // Try taking 1 digit
    const oneDigit = Number(str[index]);
    if (oneDigit in squares) {
      backtrack(index + 1, current + squares[oneDigit]);
    }
    
    // Try taking 2 digits (if available)
    if (index + 1 < str.length) {
      const twoDigits = Number(str.substring(index, index + 2));
      if (twoDigits in squares) {
        backtrack(index + 2, current + squares[twoDigits]);
      }
    }
  }
  
  backtrack(0, '');
  
  if (results.length === 0) {
    return null;
  }
  
  return Math.min(...results);
}