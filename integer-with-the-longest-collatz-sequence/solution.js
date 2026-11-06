function longestCollatz(arr) {
  function collatzLength(n) {
    let length = 0;
    while (n !== 1) {
      n = n % 2 === 0 ? n / 2 : 3 * n + 1;
      length++;
    }
    return length;
  }
  
  let maxLength = 0;
  let result;
  
  for (let num of arr) {
    const length = collatzLength(num);
    if (length > maxLength) {
      maxLength = length;
      result = num;
    }
  }
  
  return result;
}