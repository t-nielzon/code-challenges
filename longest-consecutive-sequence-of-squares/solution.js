function longestConsecutiveSquares(n) {
  let longest = [];
  
  // Try each possible starting number
  for (let start = 1; start * start <= n; start++) {
    let sum = 0;
    
    // Try extending the sequence from this start
    for (let current = start; sum < n; current++) {
      sum += current * current;
      
      // If we found a match
      if (sum === n) {
        const sequence = [];
        for (let i = start; i <= current; i++) {
          sequence.push(i);
        }
        // Keep the longest sequence
        if (sequence.length > longest.length) {
          longest = sequence;
        }
      }
      
      // Stop if sum exceeds n
      if (sum > n) break;
    }
  }
  
  return longest;
}