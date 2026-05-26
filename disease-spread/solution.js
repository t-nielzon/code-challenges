function epidemic(tm, n, s0, i0, b, a) {
  const dt = tm / n;
  let s = s0, i = i0, maxI = i0;
  for (let k = 0; k < n; k++) {
    const sNext = s - dt * b * s * i;
    const iNext = i + dt * (b * s * i - a * i);
    s = sNext;
    i = iNext;
    if (i > maxI) maxI = i;
  }
  return Math.trunc(maxI);
}