function bouncyRatio(ratio) {
  if (ratio < 0 || ratio > 99) {
    throw new Error('Ratio must be between 0 and 99');
  }
  
  function isBouncy(n) {
    const digits = String(n).split('').map(Number);
    
    let isIncreasing = true;
    let isDecreasing = true;
    
    for (let i = 1; i < digits.length; i++) {
      if (digits[i] < digits[i - 1]) {
        isIncreasing = false;
      }
      if (digits[i] > digits[i - 1]) {
        isDecreasing = false;
      }
    }
    
    return !isIncreasing && !isDecreasing;
  }
  
  let bouncyCount = 0;
  
  for (let n = 1; ; n++) {
    if (isBouncy(n)) {
      bouncyCount++;
    }
    
    const currentRatio = (bouncyCount / n) * 100;
    if (currentRatio >= ratio) {
      return n;
    }
  }
}