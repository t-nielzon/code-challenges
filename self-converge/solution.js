function selfConverge(n) {
  const width = String(n).length;

  function helper(current, seen, count) {
    const digits = String(current).padStart(width, '0');
    const desc = digits.split('').sort((a, b) => b - a).join('');
    const asc = digits.split('').sort((a, b) => a - b).join('');
    const nseq = Number(desc) - Number(asc);

    if (nseq === 0) return -1;

    count++;

    if (seen.has(nseq)) return count;

    seen.add(nseq);
    return helper(nseq, seen, count);
  }

  return helper(n, new Set(), 0);
}