function chooseBestSum(t, k, ls) {
  function helper(index, limit, needed) {
    if (needed === 0) return 0;
    if (index >= ls.length || needed > ls.length - index) return null;
    
    let best = null;
    
    // try including current element
    if (ls[index] <= limit) {
      const sum = helper(index + 1, limit - ls[index], needed - 1);
      if (sum !== null) {
        best = ls[index] + sum;
      }
    }
    
    // try excluding current element
    const sum = helper(index + 1, limit, needed);
    if (sum !== null && (best === null || sum > best)) {
      best = sum;
    }
    
    return best;
  }
  
  return helper(0, t, k);
}