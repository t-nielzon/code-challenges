function landPerimeter(arr) {
  let perimeter = 0;
  
  for (let row = 0; row < arr.length; row++) {
    for (let col = 0; col < arr[row].length; col++) {
      if (arr[row][col] === 'X') {
        // Each land tile starts with 4 sides
        let sides = 4;
        
        // Check top neighbor
        if (row > 0 && arr[row - 1][col] === 'X') {
          sides--;
        }
        
        // Check bottom neighbor
        if (row < arr.length - 1 && arr[row + 1][col] === 'X') {
          sides--;
        }
        
        // Check left neighbor
        if (col > 0 && arr[row][col - 1] === 'X') {
          sides--;
        }
        
        // Check right neighbor
        if (col < arr[row].length - 1 && arr[row][col + 1] === 'X') {
          sides--;
        }
        
        perimeter += sides;
      }
    }
  }
  
  return `Total land perimeter: ${perimeter}`;
}