function teknonymize(tree) {
  if (!tree) return;

  function getDeepestEldest(node, depth) {
    if (node.children.length === 0) {
      return { name: node.name, depth, dob: node.dateOfBirth };
    }
    let best = null;
    for (const child of node.children) {
      const candidate = getDeepestEldest(child, depth + 1);
      if (!best || candidate.depth > best.depth ||
          (candidate.depth === best.depth && candidate.dob < best.dob)) {
        best = candidate;
      }
    }
    return best;
  }

  function buildTeknonym(sex, distance, descendantName) {
    const base = sex === 'm' ? 'father' : 'mother';
    if (distance === 1) return `${base} of ${descendantName}`;
    if (distance === 2) return `grand${base} of ${descendantName}`;
    return 'great-'.repeat(distance - 2) + `grand${base} of ${descendantName}`;
  }

  function process(node) {
    if (node.children.length > 0) {
      const best = getDeepestEldest(node, 0);
      node.teknonym = buildTeknonym(node.sex, best.depth, best.name);
    }
    for (const child of node.children) {
      process(child);
    }
  }

  process(tree);
}