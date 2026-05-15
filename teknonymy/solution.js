function teknonymy(person) {
  function deepest(node) {
    if (!node.children || node.children.length === 0) {
      return { depth: 0, elder: null };
    }
    let bestDepth = 0;
    let bestElder = null;
    for (const child of node.children) {
      const sub = deepest(child);
      const candidateDepth = sub.depth + 1;
      const candidateElder = sub.elder || child;
      if (candidateDepth > bestDepth) {
        bestDepth = candidateDepth;
        bestElder = candidateElder;
      } else if (candidateDepth === bestDepth) {
        if (candidateElder.dateOfBirth < bestElder.dateOfBirth) {
          bestElder = candidateElder;
        }
      }
    }
    return { depth: bestDepth, elder: bestElder };
  }

  function assign(node) {
    const { depth, elder } = deepest(node);
    if (depth > 0 && elder) {
      const parentWord = node.sex === 'm' ? 'father' : 'mother';
      let prefix;
      if (depth === 1) prefix = '';
      else if (depth === 2) prefix = 'grand';
      else prefix = 'great-'.repeat(depth - 2) + 'grand';
      node.teknonym = prefix + parentWord + ' of ' + elder.name;
    }
    if (node.children) {
      for (const child of node.children) assign(child);
    }
  }

  assign(person);
}