function convertFracs(lst) {
  if (lst.length === 0) return [];
  
  function gcd(a, b) {
    while (b !== 0) {
      let temp = b;
      b = a % b;
      a = temp;
    }
    return a;
  }
  
  function lcm(a, b) {
    return (a * b) / gcd(a, b);
  }
  
  // Simplify input fractions first
  let simplified = lst.map(([num, denom]) => {
    let g = gcd(num, denom);
    return [num / g, denom / g];
  });
  
  // Find LCM of all denominators
  let commonDenom = simplified[0][1];
  for (let i = 1; i < simplified.length; i++) {
    commonDenom = lcm(commonDenom, simplified[i][1]);
  }
  
  // Convert each fraction to common denominator
  let result = simplified.map(([num, denom]) => {
    let newNum = num * (commonDenom / denom);
    return [newNum, commonDenom];
  });
  
  return result;
}