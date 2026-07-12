function identifyBB(bearings, weigh) {
  const sample = [];
  
  // Take i+1 bearings from each box at index i
  for (let i = 0; i < bearings.length; i++) {
    for (let j = 0; j < i + 1; j++) {
      sample.push(bearings[i]);
    }
  }
  
  // Weigh all sampled bearings at once
  const weight = weigh(...sample);
  
  // Calculate expected weight if all were normal (10g each)
  const n = bearings.length;
  const expectedWeight = 10 * (n * (n + 1) / 2);
  
  // The difference tells us which box index has the deluxe bearing
  const deluxeIndex = weight - expectedWeight;
  
  return bearings[deluxeIndex];
}