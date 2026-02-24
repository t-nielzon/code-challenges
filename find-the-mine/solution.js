function mineLocation(field) {
  for (let r = 0; r < field.length; r++) {
    const c = field[r].indexOf(1);
    if (c !== -1) return [r, c];
  }
}

module.exports = mineLocation;