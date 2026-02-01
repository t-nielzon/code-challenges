function powerSumDigTerm(n) {
  const sumOfDigits = (num) => {
    let sum = 0;
    while (num > 0) {
      sum += num % 10;
      num = Math.floor(num / 10);
    }
    return sum;
  };

  const isPowerOfDigitSum = (num) => {
    const digitSum = sumOfDigits(num);
    if (digitSum < 2) return false;
    
    let power = digitSum;
    while (power < num) {
      power *= digitSum;
    }
    return power === num;
  };

  const results = [];
  let current = 10;
  
  while (results.length < n) {
    if (isPowerOfDigitSum(current)) {
      results.push(current);
    }
    current++;
  }
  
  return results[n - 1];
}