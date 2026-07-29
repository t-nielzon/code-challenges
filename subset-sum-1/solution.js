function subsetSum(list, target) {
  function backtrack(index, remaining, currentSubset) {
    if (remaining === 0) {
      return currentSubset.slice();
    }
    
    if (index === list.length || remaining < 0) {
      return null;
    }
    
    // Try including the current element
    currentSubset.push(list[index]);
    let result = backtrack(index + 1, remaining - list[index], currentSubset);
    if (result !== null) {
      return result;
    }
    currentSubset.pop();
    
    // Try not including the current element
    return backtrack(index + 1, remaining, currentSubset);
  }
  
  return backtrack(0, target, []);
}