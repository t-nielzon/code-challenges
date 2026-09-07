function getHand(word) {
  const leftHand = new Set(['q', 'w', 'e', 'r', 't', 'a', 's', 'd', 'f', 'g', 'z', 'x', 'c', 'v', 'b']);
  const rightHand = new Set(['y', 'u', 'i', 'o', 'p', 'h', 'j', 'k', 'l', 'n', 'm']);
  
  let hasLeft = false;
  let hasRight = false;
  
  for (const letter of word) {
    if (leftHand.has(letter)) {
      hasLeft = true;
    } else if (rightHand.has(letter)) {
      hasRight = true;
    }
  }
  
  if (!hasLeft && !hasRight) {
    return 'NONE';
  } else if (hasLeft && !hasRight) {
    return 'LEFT';
  } else if (!hasLeft && hasRight) {
    return 'RIGHT';
  } else {
    return 'BOTH';
  }
}