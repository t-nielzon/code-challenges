function luckCheck(str) {
  if (typeof str !== 'string' || str.length === 0 || !/^\d+$/.test(str)) {
    throw new Error('Invalid input: expected a non-empty string of digits');
  }

  const half = Math.floor(str.length / 2);
  const left = str.slice(0, half);
  const right = str.slice(str.length - half);

  const sumDigits = s => [...s].reduce((acc, d) => acc + Number(d), 0);

  return sumDigits(left) === sumDigits(right);
}