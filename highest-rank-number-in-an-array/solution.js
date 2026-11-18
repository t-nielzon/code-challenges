function highestRankNumber(arr) {
  const freq = {};
  
  for (const num of arr) {
    freq[num] = (freq[num] || 0) + 1;
  }
  
  const maxFreq = Math.max(...Object.values(freq));
  
  return Math.max(...Object.entries(freq)
    .filter(([, count]) => count === maxFreq)
    .map(([num]) => Number(num)));
}