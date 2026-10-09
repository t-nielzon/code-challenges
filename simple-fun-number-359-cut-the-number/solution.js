function maxSum(number, target) {
  const n = number.length;
  
  // dp[i] = set of all possible sums using first i characters that don't exceed target
  const dp = Array.from({ length: n + 1 }, () => new Set());
  dp[0].add(0);
  
  for (let i = 1; i <= n; i++) {
    // try all possible positions to cut before position i
    for (let j = 0; j < i; j++) {
      const piece = parseInt(number.slice(j, i));
      for (const sum of dp[j]) {
        const newSum = sum + piece;
        if (newSum <= target) {
          dp[i].add(newSum);
        }
      }
    }
  }
  
  if (dp[n].size === 0) {
    return -1;
  }
  
  return Math.max(...dp[n]);
}