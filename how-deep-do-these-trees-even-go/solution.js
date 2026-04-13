function treeDepth(arr) {
  const depths = [];

  function collect(node, depth) {
    for (const item of node) {
      if (Array.isArray(item)) {
        collect(item, depth + 1);
      } else {
        depths.push(depth);
      }
    }
  }

  collect(arr, 1);

  if (depths.length === 0) {
    return { min: 0, max: 0, mean: 0 };
  }

  return {
    min: Math.min(...depths),
    max: Math.max(...depths),
    mean: depths.reduce((sum, d) => sum + d, 0) / depths.length
  };
}

module.exports = { treeDepth };