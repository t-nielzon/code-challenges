function findNumber(sums) {
  if (sums.length === 0) {
    return 0;
  }
  
  // Calculate number of digits from the number of sums
  let numDigits = 2;
  while (numDigits * (numDigits - 1) / 2 < sums.length) {
    numDigits++;
  }
  
  // Try each possible first digit
  const startDigit = numDigits > 1 ? 1 : 0;
  for (let firstDigit = startDigit; firstDigit <= 9; firstDigit++) {
    const digits = [firstDigit];
    let valid = true;
    
    // Derive remaining digits from first digit and sums
    for (let i = 1; i < numDigits; i++) {
      const digit = sums[i - 1] - firstDigit;
      if (digit < 0 || digit > 9) {
        valid = false;
        break;
      }
      digits.push(digit);
    }
    
    if (!valid) continue;
    
    // Verify the solution
    let sumIndex = 0;
    for (let i = 0; i < numDigits; i++) {
      for (let j = i + 1; j < numDigits; j++) {
        if (digits[i] + digits[j] !== sums[sumIndex++]) {
          valid = false;
          break;
        }
      }
      if (!valid) break;
    }
    
    if (valid) {
      return parseInt(digits.join(''));
    }
  }
  
  return 0;
}