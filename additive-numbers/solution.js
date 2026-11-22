function findAdditiveNumbers(num) {
  const n = num.length;
  
  for (let i = 1; i <= n - 2; i++) {
    for (let j = 1; i + j <= n - 1; j++) {
      if ((num[0] === '0' && i > 1) || (num[i] === '0' && j > 1)) {
        continue;
      }
      
      const result = trySequence(num, i, j);
      if (result.length > 0) {
        return result;
      }
    }
  }
  
  return [];
}

function trySequence(num, i, j) {
  let first = num.substring(0, i);
  let second = num.substring(i, i + j);
  const sequence = [first, second];
  let pos = i + j;
  
  while (pos < num.length) {
    const sum = addStrings(first, second);
    
    if (!num.startsWith(sum, pos)) {
      return [];
    }
    
    if (sum.length > 1 && sum[0] === '0') {
      return [];
    }
    
    sequence.push(sum);
    first = second;
    second = sum;
    pos += sum.length;
  }
  
  return sequence;
}

function addStrings(num1, num2) {
  let result = '';
  let carry = 0;
  let i = num1.length - 1;
  let j = num2.length - 1;
  
  while (i >= 0 || j >= 0 || carry > 0) {
    const digit1 = i >= 0 ? parseInt(num1[i]) : 0;
    const digit2 = j >= 0 ? parseInt(num2[j]) : 0;
    const sum = digit1 + digit2 + carry;
    result = (sum % 10) + result;
    carry = Math.floor(sum / 10);
    i--;
    j--;
  }
  
  return result;
}