function findSpecialIdx(vms) {
  for (let weighingIdx = 0; weighingIdx < 10; weighingIdx++) {
    // Test machines [10*weighingIdx, 10*weighingIdx+9]
    for (let j = 0; j < 10; j++) {
      const machineIdx = 10 * weighingIdx + j;
      vms[machineIdx].vend(j + 1);
    }
    
    const weight = vms.weigh();
    const expectedWeight = 55 * 100; // (1+2+...+10) * 100
    const delta = weight - expectedWeight;
    
    if (delta > 0) {
      return 10 * weighingIdx + (delta - 1);
    }
  }
}