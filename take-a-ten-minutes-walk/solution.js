function isValidWalk(walk) {
  if (walk.length !== 10) return false;
  let x = 0, y = 0;
  for (const d of walk) {
    if (d === 'n') y++;
    else if (d === 's') y--;
    else if (d === 'e') x++;
    else if (d === 'w') x--;
  }
  return x === 0 && y === 0;
}

module.exports = isValidWalk;