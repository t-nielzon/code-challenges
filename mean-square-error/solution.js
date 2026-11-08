function meanSquareError(array1, array2) {
  const sumOfSquares = array1.reduce((sum, val, i) => {
    const diff = val - array2[i];
    return sum + diff * diff;
  }, 0);
  return sumOfSquares / array1.length;
}