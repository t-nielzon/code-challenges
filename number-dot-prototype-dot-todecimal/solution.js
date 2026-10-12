Number.prototype.toDecimal = function() {
  let str = this.toString();
  
  // If not in scientific notation, return as is
  if (!/e/i.test(str)) {
    return str;
  }
  
  // Parse scientific notation
  const [coefficient, exponentStr] = str.toLowerCase().split('e');
  const exponent = parseInt(exponentStr, 10);
  
  // Extract sign and work with absolute value
  const isNegative = coefficient[0] === '-';
  const absCoefficient = isNegative ? coefficient.slice(1) : coefficient;
  const [intPart, decPart = ''] = absCoefficient.split('.');
  
  // Combine all digits
  const allDigits = intPart + decPart;
  
  // Calculate new decimal position
  const oldDecimalPos = intPart.length;
  const newDecimalPos = oldDecimalPos + exponent;
  
  let result;
  
  if (newDecimalPos <= 0) {
    // Need leading zeros
    result = '0.' + '0'.repeat(-newDecimalPos) + allDigits;
  } else if (newDecimalPos >= allDigits.length) {
    // Need trailing zeros
    result = allDigits + '0'.repeat(newDecimalPos - allDigits.length);
  } else {
    // Decimal point is within the digits
    result = allDigits.slice(0, newDecimalPos) + '.' + allDigits.slice(newDecimalPos);
  }
  
  // Add back negative sign if needed
  if (isNegative) {
    result = '-' + result;
  }
  
  return result;
};