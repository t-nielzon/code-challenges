function maxEarnings(earnings, k) {
  if (earnings.length === 0) return 0;
  
  const n = earnings.length;
  // dp[i][j] = max earnings from first i days
  // j = 0: last day was a break (or at start)
  // j = 1..k: last j consecutive days were worked
  const dp = Array(n + 1).fill(null).map(() => Array(k + 1).fill(0));
  
  for (let i = 0; i < n; i++) {
    for (let j = 0; j <= k; j++) {
      // Option 1: Skip day i (always possible, mandatory if j === k)
      dp[i + 1][0] = Math.max(dp[i + 1][0], dp[i][j]);
      
      // Option 2: Work day i (only if j < k, because after k we must break)
      if (j < k) {
        dp[i + 1][j + 1] = Math.max(dp[i + 1][j + 1], dp[i][j] + earnings[i]);
      }
    }
  }
  
  return Math.max(...dp[n]);
}