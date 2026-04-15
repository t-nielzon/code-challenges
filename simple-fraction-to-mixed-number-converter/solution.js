function mixedFraction(s) {
  const [numStr, denStr] = s.split('/');
  let num = Number(numStr);
  let den = Number(denStr);

  if (den === 0) throw new Error('Zero division error');

  const negative = (num < 0) !== (den < 0);
  num = Math.abs(num);
  den = Math.abs(den);

  const whole = Math.floor(num / den);
  let remainder = num % den;

  const g = gcd(remainder, den);
  remainder = remainder / g;
  const reducedDen = den / g;

  let result = '';
  if (whole > 0 && remainder > 0) {
    result = `${whole} ${remainder}/${reducedDen}`;
  } else if (remainder > 0) {
    result = `${remainder}/${reducedDen}`;
  } else {
    result = `${whole}`;
  }

  if (negative && (whole > 0 || remainder > 0)) {
    result = '-' + result;
  }

  return result;
}

function gcd(a, b) {
  while (b) {
    [a, b] = [b, a % b];
  }
  return a;
}