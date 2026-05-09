function powerset(nums) {
  const result = [];
  const dfs = (i, current) => {
    if (i === nums.length) {
      result.push(current.slice());
      return;
    }
    dfs(i + 1, current);
    current.push(nums[i]);
    dfs(i + 1, current);
    current.pop();
  };
  dfs(0, []);
  return result;
}