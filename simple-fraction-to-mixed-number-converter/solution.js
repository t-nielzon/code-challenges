function convertToMixedNumber(input) {
  const [x, y] = input.split('/').map(Number);

  if (y === 0) {
    throw new Error('Division by zero');
  }

  const negative = (x < 0) !== (y < 0);
  let num = Math.abs(x);
  const den = Math.abs(y);

  const gcd = (a, b) => (b === 0 ? a : gcd(b, a % b));

  const whole = Math.floor(num / den);
  let rem = num % den;

  const sign = negative && num !== 0 ? '-' : '';

  if (rem === 0) {
    return sign + whole;
  }

  const g = gcd(rem, den);
  const frac = `${rem / g}/${den / g}`;

  if (whole === 0) {
    return sign + frac;
  }

  return `${sign}${whole} ${frac}`;
}