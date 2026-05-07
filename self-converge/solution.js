function selfConverge(n) {
  const width = String(n).length;

  const step = (num) => {
    const digits = String(num).padStart(width, '0').split('');
    const desc = parseInt([...digits].sort((a, b) => b - a).join(''), 10);
    const asc = parseInt([...digits].sort((a, b) => a - b).join(''), 10);
    return desc - asc;
  };

  const recurse = (current, seen, count) => {
    const next = step(current);
    if (next === 0) return -1;
    if (seen.has(next)) return count + 1;
    seen.add(next);
    return recurse(next, seen, count + 1);
  };

  return recurse(n, new Set([n]), 0);
}