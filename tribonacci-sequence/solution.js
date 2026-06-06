function tribonacci(signature, n) {
  const result = signature.slice(0, n);

  for (let i = 3; i < n; i++) {
    result.push(result[i - 1] + result[i - 2] + result[i - 3]);
  }

  return result;
}