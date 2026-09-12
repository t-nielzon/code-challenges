function solve(m) {
  // The infinite series U(∞, x) = Σ(n=1 to ∞) n·x^n = x/(1-x)²
  // Solving x/(1-x)² = m gives the quadratic equation:
  // mx² - (2m + 1)x + m = 0
  // Using quadratic formula and selecting the root where 0 < x < 1:
  const discriminant = 4 * m + 1;
  return ((2 * m + 1) - Math.sqrt(discriminant)) / (2 * m);
}