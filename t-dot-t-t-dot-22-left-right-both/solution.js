function typeCheck(str) {
  const leftKeys = new Set('12345qwertasdfgzxcvb');
  const rightKeys = new Set('67890yuiophjkl;\'nm,./');
  
  const chars = str.toLowerCase().split('').filter(c => c !== ' ');
  
  if (chars.length === 0) {
    return '';
  }
  
  let hasLeft = false;
  let hasRight = false;
  
  for (const char of chars) {
    if (leftKeys.has(char)) {
      hasLeft = true;
    } else if (rightKeys.has(char)) {
      hasRight = true;
    }
  }
  
  if (hasLeft && hasRight) {
    return 'Both';
  } else if (hasLeft) {
    return 'Left';
  } else if (hasRight) {
    return 'Right';
  }
  
  return '';
}