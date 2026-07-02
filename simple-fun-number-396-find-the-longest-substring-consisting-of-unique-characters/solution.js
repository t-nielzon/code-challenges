function findLongestSubstring(s) {
  // sliding window: track the last index seen for each char code.
  // when a repeat is found inside the window, advance the left edge past it.
  const lastSeen = new Int32Array(65536).fill(-1);
  let start = 0;
  let best = 0;

  for (let i = 0; i < s.length; i++) {
    const code = s.charCodeAt(i);
    const prev = lastSeen[code];
    if (prev >= start) {
      start = prev + 1;
    }
    lastSeen[code] = i;
    const len = i - start + 1;
    if (len > best) best = len;
  }

  return best;
}