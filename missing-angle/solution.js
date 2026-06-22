function missingAngle(h, a, o) {
  const toDeg = (rad) => rad * 180 / Math.PI;

  if (h === 0) return toDeg(Math.atan(o / a));
  if (a === 0) return toDeg(Math.asin(o / h));
  return toDeg(Math.acos(a / h));
}