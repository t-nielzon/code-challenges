function getDartboardScore(x, y) {
  const distance = Math.sqrt(x * x + y * y);

  const BULLS_EYE_RADIUS = 12.70 / 2;
  const BULL_RADIUS = 31.8 / 2;
  const TRIPLE_INNER_RADIUS = 198 / 2;
  const TRIPLE_OUTER_RADIUS = 214 / 2;
  const DOUBLE_INNER_RADIUS = 324 / 2;
  const DOUBLE_OUTER_RADIUS = 340 / 2;

  if (distance <= BULLS_EYE_RADIUS) return "DB";
  if (distance <= BULL_RADIUS) return "SB";
  if (distance > DOUBLE_OUTER_RADIUS) return "X";

  // sectors are arranged clockwise starting from the top
  // each sector spans 18 degrees (360 / 20)
  // the order clockwise from the top: 20, 1, 18, 4, 13, 6, 10, 15, 2, 17, 3, 19, 7, 16, 8, 11, 14, 9, 12, 5
  const sectors = [20, 1, 18, 4, 13, 6, 10, 15, 2, 17, 3, 19, 7, 16, 8, 11, 14, 9, 12, 5];

  // angle from positive y-axis going clockwise
  let angle = Math.atan2(x, y);
  if (angle < 0) angle += 2 * Math.PI;

  // each sector is 18 degrees wide, offset by 9 degrees (half sector)
  // so sector 20 spans from -9 to +9 degrees from the top
  const sectorWidth = (2 * Math.PI) / 20;
  const offset = sectorWidth / 2;

  let sectorIndex = Math.floor((angle + offset) / sectorWidth) % 20;
  const number = sectors[sectorIndex];

  let prefix = "";
  if (distance >= TRIPLE_INNER_RADIUS && distance <= TRIPLE_OUTER_RADIUS) {
    prefix = "T";
  } else if (distance >= DOUBLE_INNER_RADIUS && distance <= DOUBLE_OUTER_RADIUS) {
    prefix = "D";
  }

  return prefix + number;
}

module.exports = getDartboardScore;