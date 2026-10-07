function largestPalindromicProduct(lower, upper) {
  function isPalindrome(num) {
    const str = String(num);
    return str === str.split('').reverse().join('');
  }
  
  let largest = NaN;
  
  for (let i = upper; i >= lower; i--) {
    // if max product with i can't beat current best, stop outer loop
    if (!isNaN(largest) && i * upper < largest) {
      break;
    }
    
    for (let j = upper; j >= lower; j--) {
      const product = i * j;
      
      // if product can't beat current best, stop inner loop
      if (!isNaN(largest) && product < largest) {
        break;
      }
      
      if (isPalindrome(product)) {
        if (isNaN(largest) || product > largest) {
          largest = product;
        }
      }
    }
  }
  
  return largest;
}