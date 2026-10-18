function solve(m) {
  // The infinite series U(∞, x) = x + 2x² + 3x³ + ... = Σ n*x^n
  // This converges to x/(1-x)² for |x| < 1
  // 
  // Given m = x/(1-x)², solve for x:
  // x = m(1-x)²
  // mx² - (2m+1)x + m = 0
  // 
  // Using quadratic formula, take the smaller root to ensure 0 < x < 1:
  // x = [(2m + 1) - √(4m + 1)] / (2m)
  
  return ((2 * m + 1) - Math.sqrt(4 * m + 1)) / (2 * m);
}