function Dictionary(words) {
  this.words = words;
}

Dictionary.prototype.findMostSimilar = function(term) {
  let minDistance = Infinity;
  let mostSimilar = null;
  
  for (const word of this.words) {
    const distance = levenshteinDistance(term, word);
    if (distance < minDistance) {
      minDistance = distance;
      mostSimilar = word;
    }
  }
  
  return mostSimilar;
};

function levenshteinDistance(s1, s2) {
  const m = s1.length;
  const n = s2.length;
  
  // Create a 2D array for dynamic programming
  const dp = Array(m + 1).fill(null).map(() => Array(n + 1).fill(0));
  
  // Initialize base cases
  for (let i = 0; i <= m; i++) {
    dp[i][0] = i;
  }
  for (let j = 0; j <= n; j++) {
    dp[0][j] = j;
  }
  
  // Fill in the rest of the matrix
  for (let i = 1; i <= m; i++) {
    for (let j = 1; j <= n; j++) {
      if (s1[i - 1] === s2[j - 1]) {
        dp[i][j] = dp[i - 1][j - 1];
      } else {
        dp[i][j] = 1 + Math.min(
          dp[i - 1][j],     // deletion
          dp[i][j - 1],     // insertion
          dp[i - 1][j - 1]  // replacement
        );
      }
    }
  }
  
  return dp[m][n];
}