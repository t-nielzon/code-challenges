function isPP(n) {
  for (let m = 2; m * m <= n; m++) {
    let value = m * m;
    let k = 2;
    while (value <= n) {
      if (value === n) return [m, k];
      value *= m;
      k++;
    }
  }
  return null;
}