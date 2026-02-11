function getDartboardScore(x, y) {
  // sector numbers in clockwise order starting from the top (positive y-axis)
  // each sector spans 18 degrees; sector 20 is centered at 90 degrees (top)
  const sectors = [20, 1, 18, 4, 13, 6, 10, 15, 2, 17, 3, 19, 7, 16, 8, 11, 14, 9, 12, 5];

  const distance = Math.sqrt(x * x + y * y);

  // radii (diameters / 2)
  const bullseyeRadius = 12.70 / 2;
  const bullRadius = 31.8 / 2;
  const tripleInner = 198 / 2;
  const tripleOuter = 214 / 2;
  const doubleInner = 324 / 2;
  const doubleOuter = 340 / 2;

  if (distance > doubleOuter) return "X";
  if (distance <= bullseyeRadius) return "DB";
  if (distance <= bullRadius) return "SB";

  // determine which sector the dart landed in
  // atan2 gives angle from positive x-axis, counterclockwise
  // convert so 0 degrees is at the top (positive y-axis), clockwise
  let angle = Math.atan2(x, y) * (180 / Math.PI);
  if (angle < 0) angle += 360;

  // each sector is 18 degrees wide; sector 20 is centered at 0 degrees (top)
  // offset by half a sector (9 degrees) so boundaries fall between sectors
  const sectorIndex = Math.floor(((angle + 9) % 360) / 18);
  const number = sectors[sectorIndex];

  let prefix = "";
  if (distance >= tripleInner && distance <= tripleOuter) {
    prefix = "T";
  } else if (distance >= doubleInner && distance <= doubleOuter) {
    prefix = "D";
  }

  return prefix + number;
}

module.exports = getDartboardScore;