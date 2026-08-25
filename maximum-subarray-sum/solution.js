function maxSubarraySum(arr) {
  let maxCurrent = 0;
  let maxGlobal = 0;
  
  for (let i = 0; i < arr.length; i++) {
    maxCurrent = Math.max(0, maxCurrent + arr[i]);
    maxGlobal = Math.max(maxGlobal, maxCurrent);
  }
  
  return maxGlobal;
}