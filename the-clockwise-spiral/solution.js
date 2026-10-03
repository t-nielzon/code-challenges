function createSpiral(N) {
  if (!Number.isInteger(N) || N < 1) {
    return [];
  }

  const spiral = Array(N).fill(null).map(() => Array(N).fill(0));
  
  let num = 1;
  let top = 0, bottom = N - 1, left = 0, right = N - 1;
  
  while (top <= bottom && left <= right) {
    // Move right
    for (let i = left; i <= right; i++) {
      spiral[top][i] = num++;
    }
    top++;
    
    // Move down
    for (let i = top; i <= bottom; i++) {
      spiral[i][right] = num++;
    }
    right--;
    
    // Move left
    if (top <= bottom) {
      for (let i = right; i >= left; i--) {
        spiral[bottom][i] = num++;
      }
      bottom--;
    }
    
    // Move up
    if (left <= right) {
      for (let i = bottom; i >= top; i--) {
        spiral[i][left] = num++;
      }
      left++;
    }
  }
  
  return spiral;
}