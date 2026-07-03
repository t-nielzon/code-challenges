function Xbonacci(signature, n) {
  const x = signature.length;
  const result = signature.slice(0, n);

  while (result.length < n) {
    const next = result.slice(-x).reduce((sum, v) => sum + v, 0);
    result.push(next);
  }

  return result;
}