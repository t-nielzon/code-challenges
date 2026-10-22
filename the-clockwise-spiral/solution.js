function createSpiral(N) {
  if (typeof N !== 'number' || !Number.isInteger(N) || N < 1) {
    return [];
  }
  
  const matrix = Array(N).fill(null).map(() => Array(N).fill(0));
  const directions = [[0, 1], [1, 0], [0, -1], [-1, 0]];
  let dirIndex = 0;
  let row = 0, col = 0;
  
  for (let num = 1; num <= N * N; num++) {
    matrix[row][col] = num;
    
    if (num < N * N) {
      let nextRow = row + directions[dirIndex][0];
      let nextCol = col + directions[dirIndex][1];
      
      if (nextRow < 0 || nextRow >= N || nextCol < 0 || nextCol >= N || matrix[nextRow][nextCol] !== 0) {
        dirIndex = (dirIndex + 1) % 4;
        nextRow = row + directions[dirIndex][0];
        nextCol = col + directions[dirIndex][1];
      }
      
      row = nextRow;
      col = nextCol;
    }
  }
  
  return matrix;
}