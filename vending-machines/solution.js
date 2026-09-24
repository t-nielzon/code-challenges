function findSpecialIdx(vms) {
  // Vend i+1 bars from machine i (0-indexed)
  for (let i = 0; i < 100; i++) {
    for (let j = 0; j <= i; j++) {
      vms[i].vend();
    }
  }
  
  // Weigh the pile once
  const actualWeight = vms.weigh();
  
  // Expected weight if all machines normal: 100g × (1+2+...+100) = 100 × 5050 = 505000g
  // If machine j is special: actualWeight = 505000 + (j+1) extra grams
  // Therefore: j = actualWeight - 505000 - 1
  const specialIdx = actualWeight - 505000 - 1;
  
  return specialIdx;
}