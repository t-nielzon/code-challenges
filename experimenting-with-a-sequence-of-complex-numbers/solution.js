function f(x, y, eps) {
  const mod = Math.sqrt(x * x + y * y);
  if (mod >= 1) return -1;
  // |S(n,z) - z| = |z|^(n+1) < eps
  return Math.ceil(Math.log(eps) / Math.log(mod)) - 1;
}