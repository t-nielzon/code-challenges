function buddy(start, limit) {
  const sumProperDivisors = (n) => {
    let sum = 1;
    const sqrt = Math.floor(Math.sqrt(n));
    for (let i = 2; i <= sqrt; i++) {
      if (n % i === 0) {
        sum += i;
        if (i !== n / i) sum += n / i;
      }
    }
    return sum;
  };

  for (let n = start; n <= limit; n++) {
    const m = sumProperDivisors(n) - 1;
    if (m > n && sumProperDivisors(m) - 1 === n) {
      return [n, m];
    }
  }
  return "Nothing";
}