function convertFrac(arr) {
  const gcd = (a, b) => (b === 0 ? a : gcd(b, a % b));
  const lcm = (a, b) => (a / gcd(a, b)) * b;

  // simplify each fraction first to keep the common denominator minimal
  const reduced = arr.map(([n, d]) => {
    const g = gcd(n, d);
    return [n / g, d / g];
  });

  const D = reduced.reduce((acc, [, d]) => lcm(acc, d), 1);

  return reduced.map(([n, d]) => [(n * D) / d, D]);
}