function missingAngle(h, o, a) {
  const toDeg = (r) => r * 180 / Math.PI;
  if (h === 0) return toDeg(Math.atan(o / a));
  if (o === 0) return toDeg(Math.acos(a / h));
  return toDeg(Math.asin(o / h));
}