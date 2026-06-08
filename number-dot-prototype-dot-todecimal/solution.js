Number.prototype.toDecimal = function () {
  let str = this.toString();
  if (!/e/i.test(str)) return str;

  let sign = '';
  if (str[0] === '-') {
    sign = '-';
    str = str.slice(1);
  }

  const [coef, expPart] = str.split(/e/i);
  const exp = parseInt(expPart, 10);
  const [intPart, fracPart = ''] = coef.split('.');

  // all significant digits with the decimal point removed
  const digits = intPart + fracPart;
  // where the decimal point lands, measured from the start of the digits
  const pointPos = intPart.length + exp;

  let result;
  if (pointPos <= 0) {
    result = '0.' + '0'.repeat(-pointPos) + digits;
  } else if (pointPos >= digits.length) {
    result = digits + '0'.repeat(pointPos - digits.length);
  } else {
    result = digits.slice(0, pointPos) + '.' + digits.slice(pointPos);
  }

  return sign + result;
};