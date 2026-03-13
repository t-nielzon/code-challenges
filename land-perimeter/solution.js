function landPerimeter(arr) {
  let perimeter = 0;
  for (let r = 0; r < arr.length; r++) {
    for (let c = 0; c < arr[r].length; c++) {
      if (arr[r][c] === 'X') {
        // each land tile starts with 4 edges; subtract shared edges with neighbors
        perimeter += 4;
        if (r > 0 && arr[r - 1][c] === 'X') perimeter -= 2;
        if (c > 0 && arr[r][c - 1] === 'X') perimeter -= 2;
      }
    }
  }
  return `Total land perimeter: ${perimeter}`;
}

module.exports = { landPerimeter };