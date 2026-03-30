function removNb(n) {
  const sum = n * (n + 1) / 2;
  const result = [];

  for (let a = 1; a <= n; a++) {
    // a * b = sum - a - b  =>  b = (sum - a) / (a + 1)
    const b = (sum - a) / (a + 1);
    if (b <= n && b > 0 && b === Math.floor(b)) {
      result.push([a, b]);
    }
  }

  return result;
}

module.exports = { removNb };