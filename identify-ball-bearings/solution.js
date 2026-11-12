function identifyBB(bearings, weigh) {
  const toWeigh = [];
  let expectedWeight = 0;
  
  for (let i = 0; i < bearings.length; i++) {
    const count = 1 << i;  // 2^i copies of bearing i
    expectedWeight += count * 10;
    for (let j = 0; j < count; j++) {
      toWeigh.push(bearings[i]);
    }
  }
  
  const actualWeight = weigh(...toWeigh);
  const extraWeight = actualWeight - expectedWeight;
  
  // The extra weight directly indicates which bearing is deluxe
  for (let i = 0; i < bearings.length; i++) {
    if ((1 << i) === extraWeight) {
      return bearings[i];
    }
  }
}