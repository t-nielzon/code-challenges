function unsquareDigit(num) {
  const squared = String(num);
  const validSquares = ['0', '1', '4', '9', '16', '25', '36', '49', '64', '81'];
  
  function findPartitions(str, index, current) {
    if (index === str.length) {
      return [current];
    }
    
    const results = [];
    
    for (let len = 1; len <= 2; len++) {
      const chunk = str.substring(index, index + len);
      if (validSquares.includes(chunk)) {
        const digit = Math.sqrt(parseInt(chunk));
        const nextResults = findPartitions(str, index + len, current + digit);
        results.push(...nextResults);
      }
    }
    
    return results;
  }
  
  const allPartitions = findPartitions(squared, 0, '');
  
  if (allPartitions.length === 0) {
    return null;
  }
  
  return Math.min(...allPartitions.map(p => parseInt(p)));
}