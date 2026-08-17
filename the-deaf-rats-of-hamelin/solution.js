function deafRats(string) {
  const clean = string.replace(/\s/g, '');
  const [left, right] = clean.split('P');
  
  const deafLeft = (left.match(/O~/g) || []).length;
  const deafRight = (right.match(/~O/g) || []).length;
  
  return deafLeft + deafRight;
}