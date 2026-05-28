function findBall(scales) {
  const first = scales.getWeight([0, 1, 2], [3, 4, 5]);
  if (first === 0) {
    const r = scales.getWeight([6], [7]);
    return r === -1 ? 6 : 7;
  }
  const group = first === -1 ? [0, 1, 2] : [3, 4, 5];
  const r = scales.getWeight([group[0]], [group[1]]);
  if (r === -1) return group[0];
  if (r === 1) return group[1];
  return group[2];
}