/**
 * Subset Sum
 * 
 * Given a possibly empty list of strictly positive numbers and a non-negative target number,
 * return either a subset of the list summing to the target, or null if no such subset exists.
 * 
 * The subset must consist of unique (by index) list elements.
 * If a particular value occurs more than once in the input list, you can use it up to as many times as it occurs.
 * The empty subset sums to 0.
 * If multiple valid subsets exist, return any one of them.
 * 
 * @param {number[]} list - A list of strictly positive numbers
 * @param {number} target - A non-negative target number
 * @returns {number[] | null} A subset summing to target, or null if none exists
 */
function subsetSum(list, target) {
  // dp[sum] = array of indices used to reach this sum
  const dp = {};
  dp[0] = [];
  
  for (let i = 0; i < list.length; i++) {
    const num = list[i];
    const newSums = {};
    
    // try adding current element to all previously achievable sums
    for (const sumStr in dp) {
      const sum = parseInt(sumStr);
      const newSum = sum + num;
      
      // only add if within target and not already achievable
      if (newSum <= target && !(newSum in dp) && !(newSum in newSums)) {
        newSums[newSum] = [...dp[sum], i];
      }
    }
    
    Object.assign(dp, newSums);
  }
  
  // if target is achievable, reconstruct and return the subset
  if (!(target in dp)) {
    return null;
  }
  
  return dp[target].map(i => list[i]);
}