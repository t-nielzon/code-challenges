function largestPalindromicProduct(lower, upper) {
  function isPalindrome(n) {
    const str = String(n);
    return str === str.split('').reverse().join('');
  }
  
  let maxPalindrome = -1;
  
  for (let i = upper; i >= lower; i--) {
    // break outer loop if largest possible product with i is less than current max
    if (i * upper < maxPalindrome) {
      break;
    }
    
    for (let j = upper; j >= lower; j--) {
      const product = i * j;
      
      // break inner loop if product is less than current max
      if (product < maxPalindrome) {
        break;
      }
      
      if (isPalindrome(product) && product > maxPalindrome) {
        maxPalindrome = product;
      }
    }
  }
  
  return maxPalindrome === -1 ? NaN : maxPalindrome;
}