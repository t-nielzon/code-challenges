function Xbonacci(signature, n) {
  const x = signature.length;
  const result = signature.slice(0, n);
  for (let i = x; i < n; i++) {
    let sum = 0;
    for (let j = i - x; j < i; j++) sum += result[j];
    result.push(sum);
  }
  return result;
}