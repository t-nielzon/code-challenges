function Dictionary(words) {
  this.words = words;
}

Dictionary.prototype.findMostSimilar = function(term) {
  let bestWord = '';
  let bestDist = Infinity;

  for (const word of this.words) {
    const dist = levenshtein(term, word);
    if (dist < bestDist) {
      bestDist = dist;
      bestWord = word;
      if (dist === 0) return word;
    }
  }

  return bestWord;
};

function levenshtein(a, b) {
  const m = a.length, n = b.length;
  const dp = Array(n + 1);

  for (let j = 0; j <= n; j++) dp[j] = j;

  for (let i = 1; i <= m; i++) {
    let prev = dp[0];
    dp[0] = i;
    for (let j = 1; j <= n; j++) {
      const temp = dp[j];
      dp[j] = a[i - 1] === b[j - 1]
        ? prev
        : 1 + Math.min(prev, dp[j], dp[j - 1]);
      prev = temp;
    }
  }

  return dp[n];
}