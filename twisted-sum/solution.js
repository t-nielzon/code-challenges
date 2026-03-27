function twistedSum(n) {
  let sum = 0;
  for (let i = 1; i <= n; i++) {
    let num = i;
    while (num > 0) {
      sum += num % 10;
      num = Math.floor(num / 10);
    }
  }
  return sum;
}

module.exports = twistedSum;