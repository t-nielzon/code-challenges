// solution.js
function sumOfDivisors(m, n) {
  function getSumOfSquaredDivisors(num) {
    let sum = 0;
    for (let i = 1; i * i <= num; i++) {
      if (num % i === 0) {
        sum += i * i;
        if (i !== num / i) {
          sum += (num / i) * (num / i);
        }
      }
    }
    return sum;
  }
  
  function isPerfectSquare(num) {
    const sqrt = Math.sqrt(num);
    return sqrt === Math.floor(sqrt);
  }
  
  const result = [];
  for (let i = m; i <= n; i++) {
    const squaredSum = getSumOfSquaredDivisors(i);
    if (isPerfectSquare(squaredSum)) {
      result.push([i, squaredSum]);
    }
  }
  return result;
}