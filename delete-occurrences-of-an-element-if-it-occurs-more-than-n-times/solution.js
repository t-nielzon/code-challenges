function deleteNth(arr, n) {
  const counts = {};
  return arr.filter(num => {
    counts[num] = (counts[num] || 0) + 1;
    return counts[num] <= n;
  });
}