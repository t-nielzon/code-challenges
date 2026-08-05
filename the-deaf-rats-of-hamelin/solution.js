function countDeafRats(string) {
  const cleaned = string.replace(/\s/g, '');
  const [left, right] = cleaned.split('P');
  
  let deafCount = 0;
  
  // Left of P: deaf rats going left (O~)
  if (left) {
    for (let i = 0; i < left.length; i += 2) {
      if (left.substr(i, 2) === 'O~') {
        deafCount++;
      }
    }
  }
  
  // Right of P: deaf rats going right (~O)
  if (right) {
    for (let i = 0; i < right.length; i += 2) {
      if (right.substr(i, 2) === '~O') {
        deafCount++;
      }
    }
  }
  
  return deafCount;
}