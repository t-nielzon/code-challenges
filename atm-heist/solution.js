function maximumThrill(atms) {
  const n = atms.length;
  if (n === 0) return 0;
  if (n === 1) return 2 * atms[0];
  
  // Initial answer: steal from same ATM twice
  let result = 0;
  for (let i = 0; i < n; i++) {
    result = Math.max(result, 2 * atms[i]);
  }
  
  // For two different ATMs at indices i and j where i < j:
  // value = atms[i] + atms[j] + (j - i)
  //       = (atms[i] - i) + (atms[j] + j)
  // 
  // To maximize, find best combination of:
  // - maximum of (atms[i] - i) from one index
  // - maximum of (atms[j] + j) from a different index
  
  const a = atms.map((val, i) => ({val: val - i, idx: i}));
  const b = atms.map((val, i) => ({val: val + i, idx: i}));
  
  a.sort((x, y) => y.val - x.val);
  b.sort((x, y) => y.val - x.val);
  
  // Check top-2 from each to find best pairing with different indices
  for (let i = 0; i < Math.min(2, a.length); i++) {
    for (let j = 0; j < Math.min(2, b.length); j++) {
      if (a[i].idx !== b[j].idx) {
        result = Math.max(result, a[i].val + b[j].val);
      }
    }
  }
  
  return result;
}