function mixedFraction(s) {
  const [numStr, denStr] = s.split('/');
  let num = parseInt(numStr, 10);
  let den = parseInt(denStr, 10);

  if (den === 0) throw new Error('Zero division error');

  const negative = (num < 0) !== (den < 0);
  num = Math.abs(num);
  den = Math.abs(den);

  const whole = Math.floor(num / den);
  let remainder = num % den;

  const g = gcd(remainder, den);
  remainder /= g;
  const reducedDen = den / g;

  const sign = negative && (whole > 0 || remainder > 0) ? '-' : '';

  if (remainder === 0) return sign + String(whole);
  if (whole === 0) return `${sign}${remainder}/${reducedDen}`;
  return `${sign}${whole} ${remainder}/${reducedDen}`;
}

function gcd(a, b) {
  while (b) { [a, b] = [b, a % b]; }
  return a;
}