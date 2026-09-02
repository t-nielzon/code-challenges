function isBst(node) {
  const values = [];
  
  function inOrder(n) {
    if (n === null) return;
    inOrder(n.left);
    values.push(n.value);
    inOrder(n.right);
  }
  
  inOrder(node);
  
  if (values.length <= 1) return true;
  
  // Check if strictly increasing or decreasing
  const increasing = values.every((v, i) => i === 0 || v > values[i - 1]);
  const decreasing = values.every((v, i) => i === 0 || v < values[i - 1]);
  
  return increasing || decreasing;
}