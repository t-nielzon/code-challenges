function moveZeros(arr) {
  const nonZeros = arr.filter(el => el !== 0);
  const zeroCount = arr.length - nonZeros.length;
  return [...nonZeros, ...Array(zeroCount).fill(0)];
}