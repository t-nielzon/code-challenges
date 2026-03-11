function getColumnTitle(num) {
  if (num === undefined || typeof num !== 'number' || !Number.isInteger(num)) {
    throw new TypeError('num is not an integer');
  }
  if (num < 1) {
    throw new RangeError('num must be >= 1');
  }

  let result = '';
  while (num > 0) {
    num--;
    result = String.fromCharCode(65 + (num % 26)) + result;
    num = Math.floor(num / 26);
  }
  return result;
}

module.exports = { getColumnTitle };