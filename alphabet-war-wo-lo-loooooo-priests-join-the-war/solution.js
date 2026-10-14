function alphabetWar(fight) {
  const left = { w: 4, p: 3, b: 2, s: 1, t: 0 };
  const right = { m: 4, q: 3, d: 2, z: 1, j: 0 };
  
  const leftToRight = { 4: 'm', 3: 'q', 2: 'd', 1: 'z' };
  const rightToLeft = { 4: 'w', 3: 'p', 2: 'b', 1: 's' };
  
  const chars = fight.split('');
  
  const hasBothPriests = (index) => {
    const hasLeft = index > 0 && (chars[index - 1] === 't' || chars[index - 1] === 'j');
    const hasRight = index < chars.length - 1 && (chars[index + 1] === 't' || chars[index + 1] === 'j');
    return hasLeft && hasRight;
  };
  
  for (let i = 0; i < chars.length; i++) {
    if (chars[i] === 't') {
      if (i > 0 && chars[i-1] in right && chars[i-1] !== 'j' && !hasBothPriests(i-1)) {
        chars[i-1] = rightToLeft[right[chars[i-1]]];
      }
      if (i < chars.length - 1 && chars[i+1] in right && chars[i+1] !== 'j' && !hasBothPriests(i+1)) {
        chars[i+1] = rightToLeft[right[chars[i+1]]];
      }
    } else if (chars[i] === 'j') {
      if (i > 0 && chars[i-1] in left && chars[i-1] !== 't' && !hasBothPriests(i-1)) {
        chars[i-1] = leftToRight[left[chars[i-1]]];
      }
      if (i < chars.length - 1 && chars[i+1] in left && chars[i+1] !== 't' && !hasBothPriests(i+1)) {
        chars[i+1] = leftToRight[left[chars[i+1]]];
      }
    }
  }
  
  let leftScore = 0, rightScore = 0;
  
  for (const char of chars) {
    if (char in left) leftScore += left[char];
    else if (char in right) rightScore += right[char];
  }
  
  if (leftScore > rightScore) return "Left side wins!";
  if (rightScore > leftScore) return "Right side wins!";
  return "Let's fight again!";
}