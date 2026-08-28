// solution.js
function findSpecialIdx(vms) {
  // Collect candy bars: take i bars from machine i (0 from machine 0, 1 from machine 1, etc.)
  for (let i = 0; i < 100; i++) {
    for (let j = 0; j < i; j++) {
      vms[i].vend();
    }
  }
  
  // Weigh the entire pile
  const actualWeight = vms.weigh();
  
  // If all bars weighed 100g: total bars = 0+1+2+...+99 = 4950
  // Expected weight = 4950 * 100 = 495000 grams
  const expectedWeight = 4950 * 100;
  
  // The special machine with 101g bars contributes 1 extra gram per bar taken
  // Since we took i bars from machine i, the extra weight equals i
  return actualWeight - expectedWeight;
}