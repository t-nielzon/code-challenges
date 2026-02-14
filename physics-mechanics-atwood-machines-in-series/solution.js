function atwoodMachines(masses) {
  const g = 9.80665;

  // reduce from the bottom up: combine the last two masses into an effective mass
  // until only two masses remain, then compute acceleration of the first mass
  let effective = masses[masses.length - 1];

  for (let i = masses.length - 2; i >= 1; i--) {
    effective = (4 * masses[i] * effective) / (masses[i] + effective);
  }

  // two-mass acceleration formula for masses[0] vs the effective mass
  return g * (effective - masses[0]) / (masses[0] + effective);
}

module.exports = atwoodMachines;