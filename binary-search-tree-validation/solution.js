function isBST(root) {
  const arr = [];
  (function inorder(node) {
    if (!node) return;
    inorder(node.left);
    arr.push(node.value);
    inorder(node.right);
  })(root);

  if (arr.length <= 1) return true;

  const ascending = arr[1] > arr[0];
  for (let i = 1; i < arr.length; i++) {
    if (ascending ? arr[i] <= arr[i - 1] : arr[i] >= arr[i - 1]) {
      return false;
    }
  }
  return true;
}