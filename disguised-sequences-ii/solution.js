function binomial(n, k) {
  if (k > n || k < 0) return 0;
  if (k === 0 || k === n) return 1;
  k = Math.min(k, n - k);
  let result = 1;
  for (let i = 0; i < k; i++) {
    result = result * (n - i) / (i + 1);
  }
  return result;
}

function v1(n, p) {
  let sum = 0;
  for (let k = 0; k <= n; k++) {
    const sign = (k % 2 === 0) ? 1 : -1;
    const power = Math.pow(4, n - k);
    const binom = binomial(2 * n - k, k);
    sum += sign * p * power * binom;
  }
  return sum;
}

function u1(n, p) {
  let sum = 0;
  for (let k = 0; k <= n; k++) {
    const sign = (k % 2 === 0) ? 1 : -1;
    const power = Math.pow(4, n - k);
    const binom = binomial(2 * n - k + 1, k);
    sum += sign * p * power * binom;
  }
  return sum;
}

function vEff(n, p) {
  return (2 * n + 1) * p;
}

function uEff(n, p) {
  return (n + 1) * p;
}