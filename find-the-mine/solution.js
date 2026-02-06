function mineLocation(field) {
  for (let row = 0; row < field.length; row++) {
    for (let col = 0; col < field[row].length; col++) {
      if (field[row][col] === 1) {
        return [row, col];
      }
    }
  }
}