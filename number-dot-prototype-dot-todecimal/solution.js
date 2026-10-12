Number.prototype.toDecimal = function() {
  let str = this.toString();
  
  // Check if it's in scientific notation
  if (!str.includes('e') && !str.includes('E')) {
    return str;
  }
  
  // Parse scientific notation
  const match = str.match(/^([+-])?(\d+\.?\d*)[eE]([+-]?\d+)$/);
  if (!match) {
    return str;
  }
  
  const sign = match[1] || '';
  const coefficient = match[2];
  const exponent = parseInt(match[3], 10);
  
  // Remove the decimal point from coefficient
  const decimalIndex = coefficient.indexOf('.');
  let digits, originalDecimalPos;
  
  if (decimalIndex === -1) {
    digits = coefficient;
    originalDecimalPos = coefficient.length;
  } else {
    digits = coefficient.replace('.', '');
    originalDecimalPos = decimalIndex;
  }
  
  // Calculate new decimal position
  const newDecimalPos = originalDecimalPos + exponent;
  
  let result;
  
  if (newDecimalPos <= 0) {
    result = '0.' + '0'.repeat(-newDecimalPos) + digits;
  } else if (newDecimalPos >= digits.length) {
    result = digits + '0'.repeat(newDecimalPos - digits.length);
  } else {
    result = digits.substring(0, newDecimalPos) + '.' + digits.substring(newDecimalPos);
  }
  
  return sign + result;
};