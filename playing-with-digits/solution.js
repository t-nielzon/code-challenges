function digitsToSum(n, p) {
  const digits = String(n).split('').map(Number);
  let sum = 0;
  
  for (let i = 0; i < digits.length; i++) {
    sum += Math.pow(digits[i], p + i);
  }
  
  if (sum % n === 0) {
    return sum / n;
  }
  
  return -1;
}