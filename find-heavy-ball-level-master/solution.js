function findBall(scales) {
  const first = scales.getWeight([0, 1, 2], [3, 4, 5]);
  if (first === 0) {
    const second = scales.getWeight([6], [7]);
    return second === -1 ? 6 : 7;
  }
  const group = first === -1 ? [0, 1, 2] : [3, 4, 5];
  const second = scales.getWeight([group[0]], [group[1]]);
  if (second === 0) return group[2];
  return second === -1 ? group[0] : group[1];
}