function isPaperfold(arr) {
  const n = arr.length;
  
  if (n === 0) return true;
  if (arr[0] !== 1) return false;
  if (n > 1 && arr[1] !== 1) return false;
  if (n <= 2) return true;
  
  const morphism = [
    [1, 0, 0, 0], // 00
    [1, 0, 0, 1], // 01
    [1, 1, 0, 0], // 10
    [1, 1, 0, 1]  // 11
  ];
  
  let seq = [1, 1];
  
  while (seq.length < n) {
    const next = new Array(seq.length * 2);
    let idx = 0;
    for (let i = 0; i < seq.length; i += 2) {
      const pair = (seq[i] << 1) | seq[i + 1];
      const m = morphism[pair];
      next[idx++] = m[0];
      next[idx++] = m[1];
      next[idx++] = m[2];
      next[idx++] = m[3];
    }
    seq = next;
  }
  
  for (let i = 0; i < n; i++) {
    if (arr[i] !== seq[i]) return false;
  }
  return true;
}