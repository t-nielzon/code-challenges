function decompose(n) {
  const s = String(n);
  let num, den;
  if (s.includes('/')) {
    const [a, b] = s.split('/');
    num = BigInt(a);
    den = BigInt(b);
  } else if (s.includes('.')) {
    const [intPart, fracPart] = s.split('.');
    den = 10n ** BigInt(fracPart.length);
    num = BigInt(intPart) * den + BigInt(fracPart);
  } else {
    num = BigInt(s);
    den = 1n;
  }

  if (num === 0n) return [];

  const gcd = (a, b) => {
    a = a < 0n ? -a : a;
    b = b < 0n ? -b : b;
    while (b) { [a, b] = [b, a % b]; }
    return a;
  };

  const result = [];

  // handle integer part
  if (num >= den) {
    const intPart = num / den;
    result.push(intPart.toString() + "/1");
    num = num - intPart * den;
    if (num === 0n) return result;
  }

  // reduce
  let g = gcd(num, den);
  num /= g; den /= g;

  while (num !== 0n) {
    // ceil(den / num)
    let k = (den + num - 1n) / num;
    result.push("1/" + k.toString());
    // num/den - 1/k = (num*k - den) / (den*k)
    let newNum = num * k - den;
    let newDen = den * k;
    g = gcd(newNum, newDen);
    if (g !== 0n) { newNum /= g; newDen /= g; }
    num = newNum;
    den = newDen;
  }

  return result;
}