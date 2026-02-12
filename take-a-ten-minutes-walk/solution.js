function isValidWalk(walk) {
  if (walk.length !== 10) return false;

  let x = 0;
  let y = 0;

  for (const dir of walk) {
    if (dir === 'n') y++;
    else if (dir === 's') y--;
    else if (dir === 'e') x++;
    else if (dir === 'w') x--;
  }

  return x === 0 && y === 0;
}

module.exports = isValidWalk;