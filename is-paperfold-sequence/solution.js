// solution.js
function isPaperfoldSequence(arr) {
  if (arr.length === 0) return true;
  
  const morphism = [
    [1,0,0,0],  // 0,0 → 1,0,0,0
    [1,0,0,1],  // 0,1 → 1,0,0,1
    [1,1,0,0],  // 1,0 → 1,1,0,0
    [1,1,0,1]   // 1,1 → 1,1,0,1
  ];
  
  let sequence = [1, 1];
  
  while (sequence.length < arr.length) {
    const newSequence = new Array(sequence.length * 2);
    let idx = 0;
    
    for (let i = 0; i < sequence.length; i += 2) {
      const morphed = morphism[sequence[i] * 2 + sequence[i+1]];
      newSequence[idx++] = morphed[0];
      newSequence[idx++] = morphed[1];
      newSequence[idx++] = morphed[2];
      newSequence[idx++] = morphed[3];
    }
    
    sequence = newSequence;
  }
  
  for (let i = 0; i < arr.length; i++) {
    if (sequence[i] !== arr[i]) return false;
  }
  
  return true;
}