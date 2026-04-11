function common(a, b, c) {
  const countMap = (arr) => {
    const map = {};
    for (const n of arr) {
      map[n] = (map[n] || 0) + 1;
    }
    return map;
  };

  const mapA = countMap(a);
  const mapB = countMap(b);
  const mapC = countMap(c);

  let sum = 0;
  for (const key of Object.keys(mapA)) {
    const minCount = Math.min(mapA[key], mapB[key] || 0, mapC[key] || 0);
    sum += Number(key) * minCount;
  }
  return sum;
}

module.exports = { common };