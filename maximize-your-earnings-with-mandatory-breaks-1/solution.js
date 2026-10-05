function maximizeEarnings(earnings, k) {
  const n = earnings.length;
  if (n === 0) return 0;
  
  // dp[j] = max earnings with j consecutive working days ending at current day
  // j=0 means we took a break at current day
  let dp = new Array(k + 1).fill(Number.NEGATIVE_INFINITY);
  dp[0] = 0;
  dp[1] = earnings[0];
  
  for (let i = 1; i < n; i++) {
    let newDp = new Array(k + 1).fill(Number.NEGATIVE_INFINITY);
    
    // take a break on day i (can transition from any previous state)
    newDp[0] = Math.max(...dp);
    
    // work on day i
    newDp[1] = dp[0] + earnings[i];
    for (let j = 2; j <= k; j++) {
      newDp[j] = dp[j-1] + earnings[i];
    }
    
    dp = newDp;
  }
  
  return Math.max(...dp);
}