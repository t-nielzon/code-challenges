function powerset(nums) {
  const result = [];
  
  function dfs(index, current) {
    if (index === nums.length) {
      result.push([...current]);
      return;
    }
    
    // don't take current element
    dfs(index + 1, current);
    
    // take current element
    current.push(nums[index]);
    dfs(index + 1, current);
    current.pop();
  }
  
  dfs(0, []);
  return result;
}