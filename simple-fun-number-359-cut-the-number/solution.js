function cutTheNumber(number, target) {
  const n = number.length;
  let maxSum = -1;
  
  // Iterate through all 2^(n-1) possible partitions using bitmask
  // Each bit i represents whether to cut after position i
  for (let mask = 0; mask < (1 << (n - 1)); mask++) {
    let sum = 0;
    let currentPiece = number[0];
    
    for (let i = 1; i < n; i++) {
      if (mask & (1 << (i - 1))) {
        // Cut after position i-1
        sum += parseInt(currentPiece);
        currentPiece = number[i];
      } else {
        // No cut, append to current piece
        currentPiece += number[i];
      }
    }
    // Add the last piece
    sum += parseInt(currentPiece);
    
    // Track the maximum sum that doesn't exceed target
    if (sum <= target && sum > maxSum) {
      maxSum = sum;
    }
  }
  
  return maxSum;
}