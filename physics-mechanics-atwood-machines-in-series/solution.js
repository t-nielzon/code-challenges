function atwoodMachines(masses) {
  const g = 9.80665;

  // reduce all masses from the end into a single effective mass
  let m2 = masses[masses.length - 1];
  for (let i = masses.length - 2; i >= 1; i--) {
    m2 = (4 * masses[i] * m2) / (masses[i] + m2);
  }

  const m1 = masses[0];
  return g * (m2 - m1) / (m1 + m2);
}

module.exports = atwoodMachines;