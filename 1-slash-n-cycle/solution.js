function cycle(n) {
  function gcd(a, b) {
    while (b) { [a, b] = [b, a % b]; }
    return a;
  }

  if (gcd(n, 10) !== 1) return -1;

  let remainder = 1 % n;
  for (let i = 1; i <= n; i++) {
    remainder = (remainder * 10) % n;
    if (remainder === 1) return i;
  }

  return -1;
}

module.exports = { cycle };