function handType(str) {
  if (!str || str.trim() === '') return '';
  
  const left = new Set('`12345qwertyasdfgzxcvb');
  const right = new Set('67890-=yuiophjkl;\'[]\\nm,./?');
  
  const shiftMap = {
    '!': '1', '@': '2', '#': '3', '$': '4', '%': '5',
    '^': '6', '&': '7', '*': '8', '(': '9', ')': '0',
    '_': '-', '+': '=', ':': ';', '"': "'", '{': '[', '}': ']', '|': '\\',
    '<': ',', '>': '.', '?': '/', '~': '`'
  };
  
  let hasLeft = false;
  let hasRight = false;
  
  for (const char of str) {
    if (char === ' ') continue;
    
    let key = char.toLowerCase();
    
    if (shiftMap[char]) {
      key = shiftMap[char];
    }
    
    if (left.has(key)) {
      hasLeft = true;
    } else if (right.has(key)) {
      hasRight = true;
    }
  }
  
  if (hasLeft && hasRight) return 'Both';
  if (hasLeft) return 'Left';
  if (hasRight) return 'Right';
  return '';
}