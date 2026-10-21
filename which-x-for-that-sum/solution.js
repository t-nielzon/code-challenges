function solve(m) {
  // The infinite series sum is: U(∞, x) = x/(1-x)²
  // Solving m = x/(1-x)² gives: mx² - (2m+1)x + m = 0
  // Using quadratic formula: x = [(2m+1) ± √(4m + 1)] / (2m)
  // For 0 < x < 1, take the minus sign
  const discriminant = 4 * m + 1;
  return ((2 * m + 1) - Math.sqrt(discriminant)) / (2 * m);
}