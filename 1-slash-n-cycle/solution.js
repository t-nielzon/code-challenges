function cycleLength(n) {
  function gcd(a, b) {
    while (b !== 0) {
      [a, b] = [b, a % b];
    }
    return a;
  }
  
  if (gcd(n, 10) !== 1) {
    return -1;
  }
  
  let remainder = 1;
  const remainderPositions = {};
  let position = 0;
  
  while (!(remainder in remainderPositions)) {
    remainderPositions[remainder] = position;
    remainder = (remainder * 10) % n;
    position++;
  }
  
  return position - remainderPositions[remainder];
}