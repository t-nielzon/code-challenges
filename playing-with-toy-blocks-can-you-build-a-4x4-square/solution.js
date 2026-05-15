function buildSquare(blocks) {
  const c = [0, 0, 0, 0, 0];
  for (const b of blocks) c[b]++;
  const n1 = c[1], n2 = c[2], n3 = c[3], n4 = c[4];

  for (let a = 0; a <= Math.min(4, n4); a++) {
    for (let b = 0; b <= Math.min(4 - a, n3); b++) {
      for (let cc = 0; cc <= 4 - a - b; cc++) {
        for (let d = 0; d <= 4 - a - b - cc; d++) {
          const e = 4 - a - b - cc - d;
          if (e < 0) continue;
          if (2 * cc + d > n2) continue;
          if (b + 2 * d + 4 * e > n1) continue;
          return true;
        }
      }
    }
  }
  return false;
}