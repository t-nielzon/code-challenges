function unsquare(num) {
  const squared = num.toString();
  const squareToDigit = {
    '0': '0',
    '1': '1',
    '4': '2',
    '9': '3',
    '16': '4',
    '25': '5',
    '36': '6',
    '49': '7',
    '64': '8',
    '81': '9'
  };
  
  const results = [];
  
  function backtrack(index, current) {
    if (index === squared.length) {
      results.push(parseInt(current));
      return;
    }
    
    // try single character
    const oneChar = squared[index];
    if (squareToDigit[oneChar]) {
      backtrack(index + 1, current + squareToDigit[oneChar]);
    }
    
    // try two characters
    if (index + 1 < squared.length) {
      const twoChar = squared.substring(index, index + 2);
      if (squareToDigit[twoChar]) {
        backtrack(index + 2, current + squareToDigit[twoChar]);
      }
    }
  }
  
  backtrack(0, '');
  
  return results.length === 0 ? null : Math.min(...results);
}