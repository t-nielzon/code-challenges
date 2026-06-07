function sevenAte9(arr) {
  const result = [];
  for (let i = 0; i < arr.length; i++) {
    if (arr[i] === 7 && arr[i + 1] === 8 && arr[i + 2] === 9) {
      // skip the 7 here; it will be placed after the 9
      result.push(8, 9, 7);
      i += 2;
    } else {
      result.push(arr[i]);
    }
  }
  return result;
}