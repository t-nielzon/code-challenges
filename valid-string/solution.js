function validString(dictionary, word) {
  const n = word.length;
  const dp = new Array(n + 1).fill(false);
  dp[0] = true;

  for (let i = 1; i <= n; i++) {
    for (const w of dictionary) {
      const len = w.length;
      if (i >= len && dp[i - len] && word.slice(i - len, i) === w) {
        dp[i] = true;
        break;
      }
    }
  }

  return dp[n];
}