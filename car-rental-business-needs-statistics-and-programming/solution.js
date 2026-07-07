function prob_simpson(lamb, x, operator) {
  const e = 2.7182818284;
  
  function factorial(n) {
    if (n === 0 || n === 1) return 1;
    let result = 1;
    for (let i = 2; i <= n; i++) {
      result *= i;
    }
    return result;
  }
  
  function poissonProb(lambda, k) {
    return Math.pow(lambda, k) * Math.pow(e, -lambda) / factorial(k);
  }
  
  if (!operator) {
    return poissonProb(lamb, x);
  }
  
  let sum = 0;
  let limit = operator === '>' ? x - 1 : x;
  
  for (let i = 0; i <= limit; i++) {
    sum += poissonProb(lamb, i);
  }
  
  return (operator === '<' || operator === '<=') ? 1 - sum : sum;
}