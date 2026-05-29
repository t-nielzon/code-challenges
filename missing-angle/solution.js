function missingAngle(h, o, a) {
  if (h === 0) return Math.atan(o / a) * 180 / Math.PI;
  if (o === 0) return Math.acos(a / h) * 180 / Math.PI;
  return Math.asin(o / h) * 180 / Math.PI;
}