function countRectangles(points) {
  const set = new Set(points.map(([x, y]) => x + ',' + y));
  let count = 0;

  for (let i = 0; i < points.length; i++) {
    for (let j = i + 1; j < points.length; j++) {
      const [x1, y1] = points[i];
      const [x2, y2] = points[j];

      // only consider pairs that form a diagonal (different x and different y)
      if (x1 === x2 || y1 === y2) continue;

      // check if the other two corners exist
      if (set.has(x1 + ',' + y2) && set.has(x2 + ',' + y1)) {
        count++;
      }
    }
  }

  // each rectangle is counted twice (once per diagonal pair)
  return count / 2;
}