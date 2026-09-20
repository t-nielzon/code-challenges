function comp(a, b) {
  if (a == null || b == null) return false;
  if (a.length !== b.length) return false;
  
  const aSorted = a.map(x => x * x).sort((x, y) => x - y);
  const bSorted = [...b].sort((x, y) => x - y);
  
  return aSorted.every((val, idx) => val === bSorted[idx]);
}