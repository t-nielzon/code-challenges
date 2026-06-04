function solve(s) {
  // an odd-length string can never be balanced
  if (s.length % 2 !== 0) return -1;

  let open = 0;   // unmatched "(" waiting for a closer
  let close = 0;  // unmatched ")" that had no opener

  for (const c of s) {
    if (c === '(') {
      open++;
    } else if (open > 0) {
      open--;
    } else {
      close++;
    }
  }

  // each pair of like unmatched brackets needs one reversal
  return Math.ceil(open / 2) + Math.ceil(close / 2);
}