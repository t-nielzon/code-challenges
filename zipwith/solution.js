function zipWith(fn, a1, a2) {
  const length = Math.min(a1.length, a2.length);
  const result = [];
  for (let i = 0; i < length; i++) {
    result.push(fn(a1[i], a2[i]));
  }
  return result;
}