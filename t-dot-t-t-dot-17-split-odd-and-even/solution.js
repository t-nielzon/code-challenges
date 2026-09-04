function splitOddEven(n) {
  const str = String(n);
  const result = [];
  let current = str[0];
  
  for (let i = 1; i < str.length; i++) {
    const currentDigit = parseInt(str[i]);
    const lastDigit = parseInt(str[i - 1]);
    
    if ((currentDigit % 2) !== (lastDigit % 2)) {
      result.push(parseInt(current));
      current = str[i];
    } else {
      current += str[i];
    }
  }
  
  result.push(parseInt(current));
  return result;
}