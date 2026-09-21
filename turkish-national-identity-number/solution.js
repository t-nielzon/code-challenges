function isValidTurkishId(input) {
  // Check if input is a valid number
  if (typeof input !== 'number' || !Number.isInteger(input) || input < 0) {
    return false;
  }
  
  const str = String(input);
  
  // Must have exactly 11 digits
  if (str.length !== 11) {
    return false;
  }
  
  // First digit can't be zero
  if (str[0] === '0') {
    return false;
  }
  
  // Convert to array of numbers
  const digits = str.split('').map(Number);
  
  // Check 10th digit rule: (odd_sum * 7 - even_sum) % 10 = 10th digit
  const oddSum = digits[0] + digits[2] + digits[4] + digits[6] + digits[8];
  const evenSum = digits[1] + digits[3] + digits[5] + digits[7];
  if ((oddSum * 7 - evenSum) % 10 !== digits[9]) {
    return false;
  }
  
  // Check 11th digit rule: sum of first 10 digits % 10 = 11th digit
  const firstTenSum = digits.slice(0, 10).reduce((a, b) => a + b, 0);
  if (firstTenSum % 10 !== digits[10]) {
    return false;
  }
  
  return true;
}