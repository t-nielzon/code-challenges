function preOrder(node) {
  if (!node) return [];
  return [node.data, ...preOrder(node.left), ...preOrder(node.right)];
}

function inOrder(node) {
  if (!node) return [];
  return [...inOrder(node.left), node.data, ...inOrder(node.right)];
}

function postOrder(node) {
  if (!node) return [];
  return [...postOrder(node.left), ...postOrder(node.right), node.data];
}