function brainfuck(str) {
  let bf = '';
  let currentValue = 0;
  
  for (let char of str) {
    const targetValue = char.charCodeAt(0);
    const diff = targetValue - currentValue;
    
    if (diff > 0) {
      bf += '+'.repeat(diff);
    } else if (diff < 0) {
      bf += '-'.repeat(-diff);
    }
    
    bf += '.';
    currentValue = targetValue;
  }
  
  return bf;
}