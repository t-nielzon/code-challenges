function convertFrac(lst) {
  const gcd = (a, b) => b === 0n ? a : gcd(b, a % b);
  const lcm = (a, b) => a * b / gcd(a, b);

  const reduced = lst.map(([n, d]) => {
    const bn = BigInt(n), bd = BigInt(d);
    const g = gcd(bn, bd);
    return [bn / g, bd / g];
  });

  const D = reduced.reduce((acc, [, d]) => lcm(acc, d), 1n);

  return reduced.map(([n, d]) => `(${n * (D / d)},${D})`).join('');
}