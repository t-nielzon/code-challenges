function calculateRAM(E, d, data) {
  // Conversion constant: A_r = k * E * t² / d²
  // Derived from KE = (1/2)mv² with unit conversions
  // E in aJ, d in mm, t in ns, N_A = 6.022e23
  const k = 0.0012044;
  
  // Calculate apparent m/z for each measurement
  const measurements = data.map(([I, t]) => {
    const apparentMz = k * E * t * t / (d * d);
    return { apparentMz, I };
  });
  
  // Group measurements by actual isotope
  const isotopeMap = {};
  
  for (const { apparentMz, I } of measurements) {
    let actualMass;
    
    // Determine if this is from a 1+ or 2+ ion
    // 1+ ions: m/z is integer (A_r)
    // 2+ ions: m/z is X.5 (A_r/2)
    const fractionalPart = apparentMz - Math.floor(apparentMz);
    
    if (fractionalPart < 0.1 || fractionalPart > 0.9) {
      // Close to integer: 1+ ion
      actualMass = Math.round(apparentMz);
    } else {
      // Close to X.5: 2+ ion, so actual mass is double
      actualMass = Math.round(2 * apparentMz);
    }
    
    // Accumulate current for this isotope
    isotopeMap[actualMass] = (isotopeMap[actualMass] || 0) + I;
  }
  
  // Calculate relative atomic mass as weighted average
  const totalCurrent = Object.values(isotopeMap).reduce((a, b) => a + b, 0);
  let ram = 0;
  
  for (const [mass, current] of Object.entries(isotopeMap)) {
    ram += parseInt(mass) * (current / totalCurrent);
  }
  
  // Round to nearest 0.01
  return Math.round(ram * 100) / 100;
}