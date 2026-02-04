function powerset(nums) {
  const result = [];
  
  function dfs(index, current) {
    if (index === nums.length) {
      result.push([...current]);
      return;
    }
    
    // First, choose NOT to take the current element
    dfs(index + 1, current);
    
    // Then, choose to take the current element
    current.push(nums[index]);
    dfs(index + 1, current);
    current.pop();
  }
  
  dfs(0, []);
  return result;
}