function cantor(n) {
  // Find which diagonal contains the nth element
  // Diagonal d contains elements where the sum (numerator + denominator) = d + 1
  // Total elements up to diagonal d is d(d+1)/2
  const d = Math.ceil((-1 + Math.sqrt(1 + 8 * n)) / 2);
  
  // Find position within the diagonal (1-indexed)
  const position = n - d * (d - 1) / 2;
  
  // In diagonal d, numerator goes 1 to d, denominator goes d to 1
  const numerator = position;
  const denominator = d + 1 - position;
  
  return `${numerator}/${denominator}`;
}