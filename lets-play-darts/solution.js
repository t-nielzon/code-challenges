function getDartboardScore(x, y) {
  const r = Math.sqrt(x * x + y * y);
  if (r > 170) return "X";
  if (r <= 6.35) return "DB";
  if (r <= 15.9) return "SB";

  const sections = [20, 5, 12, 9, 14, 11, 8, 16, 7, 19, 3, 17, 2, 15, 10, 6, 13, 4, 18, 1];
  let angle = Math.atan2(y, x) * 180 / Math.PI;
  if (angle < 0) angle += 360;
  const idx = Math.floor(((angle - 81 + 360) % 360) / 18);
  const num = sections[idx];

  if (r <= 99) return String(num);
  if (r <= 107) return "T" + num;
  if (r <= 162) return String(num);
  return "D" + num;
}