function palindromeNum(num) {
  if (!Number.isInteger(num) || num < 0) {
    return "Not valid";
  }
  
  const str = String(num);
  const palindromes = new Set();
  
  for (let i = 0; i < str.length; i++) {
    for (let j = i + 2; j <= str.length; j++) {
      const substring = str.substring(i, j);
      
      if (substring[0] === '0' || substring[substring.length - 1] === '0') {
        continue;
      }
      
      if (substring === substring.split('').reverse().join('')) {
        palindromes.add(parseInt(substring));
      }
    }
  }
  
  if (palindromes.size === 0) {
    return "No palindromes found";
  }
  
  return Array.from(palindromes).sort((a, b) => a - b);
}