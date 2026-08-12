function maxSum(arr) {
  let maxCurrent = 0;
  let maxGlobal = 0;
  
  for (let num of arr) {
    maxCurrent = Math.max(0, maxCurrent + num);
    maxGlobal = Math.max(maxGlobal, maxCurrent);
  }
  
  return maxGlobal;
}