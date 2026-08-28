function pascalsTriangle(n) {
  const result = [];
  
  for (let row = 0; row < n; row++) {
    if (row === 0) {
      result.push(1);
    } else {
      const newRow = [1];
      
      // Start index of the previous row in the flattened array
      const prevRowStart = (row - 1) * row / 2;
      
      // Add middle elements by summing adjacent elements from previous row
      for (let col = 1; col < row; col++) {
        newRow.push(result[prevRowStart + col - 1] + result[prevRowStart + col]);
      }
      
      newRow.push(1);
      result.push(...newRow);
    }
  }
  
  return result;
}