function squares(x, y) {
  let total = 0;
  for (let k = 1; k <= y; k++) {
    total += (x - k + 1) * (y - k + 1);
  }
  return total;
}