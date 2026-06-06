function compress(length, getChar) {
  let result = '';
  let start = 0;

  while (start < length) {
    const ch = getChar(start);

    // each character occupies exactly one contiguous block, so the block's
    // last index can be found by binary searching the rightmost position
    // whose character still equals ch.
    let lo = start;
    let hi = length - 1;
    while (lo < hi) {
      const mid = lo + Math.floor((hi - lo + 1) / 2);
      if (getChar(mid) === ch) {
        lo = mid;
      } else {
        hi = mid - 1;
      }
    }

    const count = lo - start + 1;
    result += count + ch;
    start = lo + 1;
  }

  return result;
}