function decompose(nOverd) {
  let num, den;

  if (nOverd.includes("/")) {
    [num, den] = nOverd.split("/").map(BigInt);
  } else {
    // decimal input
    const parts = nOverd.split(".");
    if (parts.length === 1 || !parts[1]) {
      num = BigInt(parts[0]);
      den = 1n;
    } else {
      const decimals = parts[1].length;
      den = 10n ** BigInt(decimals);
      num = BigInt(parts[0]) * den + BigInt(parts[1]);
    }
  }

  // simplify with gcd
  function gcd(a, b) {
    a = a < 0n ? -a : a;
    b = b < 0n ? -b : b;
    while (b) {
      [a, b] = [b, a % b];
    }
    return a;
  }

  let g = gcd(num, den);
  num /= g;
  den /= g;

  if (num === 0n) return [];

  const result = [];

  // greedy algorithm (Sylvester-Fibonacci)
  while (num > 0n) {
    if (num === 1n) {
      result.push(`1/${den}`);
      break;
    }

    // ceiling of den/num
    let d = den / num;
    if (den % num !== 0n) d += 1n;

    if (d === 1n) {
      // integer part
      const intPart = num / den;
      result.push(`${intPart}`);
      num = num - intPart * den;
    } else {
      result.push(`1/${d}`);
      // num/den - 1/d = (num*d - den) / (den*d)
      num = num * d - den;
      den = den * d;
    }

    g = gcd(num, den);
    if (g > 1n) {
      num /= g;
      den /= g;
    }
  }

  return result;
}

module.exports = { decompose };