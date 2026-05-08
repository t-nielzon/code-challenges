function sumProperDivisors(n) {
  if (n < 2) return 0;
  let sum = 1;
  const sqrtN = Math.floor(Math.sqrt(n));
  for (let i = 2; i <= sqrtN; i++) {
    if (n % i === 0) {
      sum += i;
      const pair = n / i;
      if (pair !== i) sum += pair;
    }
  }
  return sum;
}

function buddy(start, limit) {
  for (let n = start; n <= limit; n++) {
    const m = sumProperDivisors(n) - 1;
    if (m > n && sumProperDivisors(m) - 1 === n) {
      return [n, m];
    }
  }
  return "Nothing";
}