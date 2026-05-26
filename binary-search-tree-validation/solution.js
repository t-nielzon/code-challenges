function isBST(node) {
  const arr = [];
  (function inorder(n) {
    if (!n) return;
    inorder(n.left);
    arr.push(n.value);
    inorder(n.right);
  })(node);

  const ascending = arr.every((v, i) => i === 0 || arr[i - 1] < v);
  const descending = arr.every((v, i) => i === 0 || arr[i - 1] > v);
  return ascending || descending;
}