function decompose(str) {
  // Parse the input
  let numerator, denominator;
  
  if (str.includes('/')) {
    [numerator, denominator] = str.split('/').map(Number);
  } else if (str.includes('.')) {
    // Convert decimal to fraction
    const decimalPart = str.split('.')[1];
    const decimalPlaces = decimalPart.length;
    numerator = parseInt(str.replace('.', ''));
    denominator = Math.pow(10, decimalPlaces);
  } else {
    // It's an integer
    numerator = parseInt(str);
    denominator = 1;
  }
  
  if (numerator === 0) return [];
  
  // Simplify the fraction using GCD
  const gcd = (a, b) => b === 0 ? a : gcd(b, a % b);
  const g = gcd(Math.abs(numerator), denominator);
  numerator = Math.abs(numerator) / g;
  denominator = denominator / g;
  
  const result = [];
  
  // Extract integer part
  if (numerator >= denominator) {
    const intPart = Math.floor(numerator / denominator);
    result.push(intPart.toString());
    numerator = numerator % denominator;
  }
  
  // Greedy Egyptian fraction decomposition
  while (numerator > 0) {
    // Find smallest k such that 1/k <= numerator/denominator
    // This means k = ceil(denominator/numerator)
    const k = Math.ceil(denominator / numerator);
    result.push(`1/${k}`);
    
    // Subtract 1/k from numerator/denominator
    // numerator/denominator - 1/k = (numerator*k - denominator) / (denominator*k)
    numerator = numerator * k - denominator;
    denominator = denominator * k;
    
    // Simplify to avoid overflow
    const g = gcd(numerator, denominator);
    numerator = numerator / g;
    denominator = denominator / g;
  }
  
  return result;
}