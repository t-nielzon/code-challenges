function deleteNth(arr, n) {
  const counts = {};
  return arr.filter(item => {
    counts[item] = (counts[item] || 0) + 1;
    return counts[item] <= n;
  });
}

module.exports = { deleteNth };