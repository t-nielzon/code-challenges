function solve(weights, N) {
  const n = weights.length;

  function search(i, left, right) {
    if (i === n) {
      const leftSum = left.reduce((a, b) => a + b, 0);
      const rightSum = right.reduce((a, b) => a + b, 0) + N;
      if (leftSum === rightSum && leftSum > 0) {
        return [left.slice(), right.slice()];
      }
      return null;
    }

    // skip this weight
    let result = search(i + 1, left, right);
    if (result) return result;

    // place on left pan
    left.push(weights[i]);
    result = search(i + 1, left, right);
    if (result) return result;
    left.pop();

    // place on right pan
    right.push(weights[i]);
    result = search(i + 1, left, right);
    if (result) return result;
    right.pop();

    return null;
  }

  return search(0, [], []) || [];
}