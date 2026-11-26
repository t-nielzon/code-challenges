function findSpecialIdx(vms) {
  // Take i bars from machine i to identify which one is special
  for (let i = 0; i < 100; i++) {
    for (let j = 0; j < i; j++) {
      vms[i].vend();
    }
  }
  
  // Weigh the pile once
  const weight = vms.weigh();
  
  // Expected weight if all bars were 100g
  // 0*100 + 1*100 + 2*100 + ... + 99*100 = 100 * sum(0 to 99) = 100 * 4950
  const expectedWeight = 100 * 4950;
  
  // The difference from expected weight tells us the special machine index
  return weight - expectedWeight;
}