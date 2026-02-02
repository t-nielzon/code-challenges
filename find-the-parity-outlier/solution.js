function findOutlier(integers) {
  const evens = [];
  const odds = [];
  
  for (const num of integers) {
    if (num % 2 === 0) {
      evens.push(num);
    } else {
      odds.push(num);
    }
    
    if (evens.length > 1 && odds.length === 1) {
      return odds[0];
    }
    if (odds.length > 1 && evens.length === 1) {
      return evens[0];
    }
  }
  
  return evens.length === 1 ? evens[0] : odds[0];
}

module.exports = { findOutlier };