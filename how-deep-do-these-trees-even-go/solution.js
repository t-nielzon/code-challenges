function deepestNest(arr) {
  const depths = [];

  function traverse(node, depth) {
    for (const el of node) {
      if (Array.isArray(el)) {
        traverse(el, depth + 1);
      } else {
        depths.push(depth);
      }
    }
  }

  traverse(arr, 0);

  if (depths.length === 0) return { min: 0, max: 0, mean: 0 };

  const min = Math.min(...depths);
  const max = Math.max(...depths);
  const mean = depths.reduce((a, b) => a + b, 0) / depths.length;

  return { min, max, mean };
}

module.exports = { deepestNest };