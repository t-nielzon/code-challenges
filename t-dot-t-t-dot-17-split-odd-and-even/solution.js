function splitOddEven(n) {
  const str = String(n);
  const result = [];
  let currentGroup = str[0];
  
  for (let i = 1; i < str.length; i++) {
    const currentDigit = parseInt(str[i]);
    const prevDigit = parseInt(str[i - 1]);
    
    if ((currentDigit % 2) === (prevDigit % 2)) {
      currentGroup += str[i];
    } else {
      result.push(parseInt(currentGroup));
      currentGroup = str[i];
    }
  }
  
  result.push(parseInt(currentGroup));
  return result;
}