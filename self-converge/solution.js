function selfConverge(n) {
  const width = String(n).length;

  function iterate(current, seen) {
    const padded = String(current).padStart(width, '0');
    const descending = Number([...padded].sort((a, b) => b - a).join(''));
    const ascending = Number([...padded].sort((a, b) => a - b).join(''));
    const nseq = descending - ascending;

    if (nseq === 0) return -1;
    if (seen.has(nseq)) return seen.size;

    seen.add(nseq);
    return iterate(nseq, seen);
  }

  return iterate(n, new Set());
}