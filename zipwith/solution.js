function zipWith(fn, array1, array2) {
  const minLength = Math.min(array1.length, array2.length);
  return Array.from({ length: minLength }, (_, i) => fn(array1[i], array2[i]));
}