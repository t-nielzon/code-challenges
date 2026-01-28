function countRectangles(points) {
  const set = new Set(points.map(p => p[0] + ',' + p[1]));
  let count = 0;

  for (let i = 0; i < points.length; i++) {
    for (let j = i + 1; j < points.length; j++) {
      const [x1, y1] = points[i];
      const [x2, y2] = points[j];

      if (x1 !== x2 && y1 !== y2) {
        if (set.has(x1 + ',' + y2) && set.has(x2 + ',' + y1)) {
          count++;
        }
      }
    }
  }

  return count / 2;
}