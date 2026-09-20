function convertFracs(fracs) {
  if (fracs.length === 0) return fracs;
  
  function gcd(a, b) {
    while (b !== 0) {
      const temp = b;
      b = a % b;
      a = temp;
    }
    return a;
  }
  
  function lcm(a, b) {
    return (a * b) / gcd(a, b);
  }
  
  // reduce each input fraction to lowest terms
  const reduced = fracs.map(frac => {
    const g = gcd(frac[0], frac[1]);
    return [frac[0] / g, frac[1] / g];
  });
  
  // find lcm of all reduced denominators
  const commonDenom = reduced.reduce((acc, frac) => lcm(acc, frac[1]), reduced[0][1]);
  
  // convert each original fraction to the common denominator
  const result = fracs.map(frac => {
    const newNumer = frac[0] * (commonDenom / frac[1]);
    return [newNumer, commonDenom];
  });
  
  return result;
}