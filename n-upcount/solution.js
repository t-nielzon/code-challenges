function upcount(arr, n) {
  let sum = 0;
  let count = 0;
  for (let i = 0; i < arr.length; i++) {
    const prev = sum;
    sum += arr[i];
    if (prev < n && sum >= n) {
      count++;
    }
  }
  return count;
}