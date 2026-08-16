function isPaperfoldSequence(arr) {
  if (arr.length === 0) return true;
  
  const morphism = [
    [[1, 0, 0, 0], [1, 0, 0, 1]],
    [[1, 1, 0, 0], [1, 1, 0, 1]]
  ];
  
  let sequence = [1, 1];
  
  while (sequence.length < arr.length) {
    const next = [];
    for (let i = 0; i < sequence.length; i += 2) {
      const replacement = morphism[sequence[i]][sequence[i + 1]];
      next.push(replacement[0], replacement[1], replacement[2], replacement[3]);
    }
    sequence = next;
  }
  
  for (let i = 0; i < arr.length; i++) {
    if (sequence[i] !== arr[i]) return false;
  }
  return true;
}