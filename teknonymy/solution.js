function teknpionymize(tree) {
  function getMaxDepthElder(node, depth) {
    if (!node.children || node.children.length === 0) {
      return { depth, name: node.name, dob: node.dateOfBirth };
    }
    let best = null;
    for (const child of node.children) {
      const result = getMaxDepthElder(child, depth + 1);
      if (!best || result.depth > best.depth || (result.depth === best.depth && result.dob < best.dob)) {
        best = result;
      }
    }
    return best;
  }

  function buildTeknonym(sex, generationDist, elderName) {
    const parent = sex === 'm' ? 'father' : 'mother';
    if (generationDist === 1) return `${parent} of ${elderName}`;
    const grandparent = sex === 'm' ? 'grandfather' : 'grandmother';
    if (generationDist === 2) return `${grandparent} of ${elderName}`;
    const greats = 'great-'.repeat(generationDist - 2);
    return `${greats}${grandparent} of ${elderName}`;
  }

  function process(node) {
    if (!node.children || node.children.length === 0) {
      node.teknonym = '';
      return;
    }
    for (const child of node.children) {
      process(child);
    }
    const result = getMaxDepthElder(node, 0);
    if (result.depth === 0) {
      node.teknonym = '';
    } else {
      node.teknonym = buildTeknonym(node.sex, result.depth, result.name);
    }
  }

  process(tree);
}