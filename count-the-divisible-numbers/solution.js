function divisibleCount(x, y, k) {
  return Math.floor(y / k) - Math.floor((x - 1) / k);
}