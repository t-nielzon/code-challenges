function moveZeros(arr) {
  const zeroCount = arr.filter(x => x === 0).length;
  const nonZeros = arr.filter(x => x !== 0);
  return nonZeros.concat(Array(zeroCount).fill(0));
}