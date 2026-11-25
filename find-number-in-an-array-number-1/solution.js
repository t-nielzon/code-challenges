function findNum(arr) {
  const n = Math.max(...arr);
  const len = arr.length;
  const expectedSum = (n * (n + 1)) / 2;
  const actualSum = arr.reduce((sum, num) => sum + num, 0);
  
  // Case 1: n+1 elements with one duplicate
  if (len === n + 1) {
    return actualSum - expectedSum;
  }
  
  // Case 2: 2*n-1 elements with one unique (others appear twice)
  return 2 * expectedSum - actualSum;
}