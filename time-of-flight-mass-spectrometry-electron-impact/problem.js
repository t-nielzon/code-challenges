// Time of Flight Mass Spectrometry: Electron Impact
//
// Calculate the relative atomic mass (RAM) of a sample using time-of-flight mass spectrometry data.
//
// Given:
// - E: kinetic energy each ion is accelerated to (in aJ = 10^-18 J)
// - d: length of flight tube (in mm = 10^-3 m)
// - data: array of [current (µA), time (ns)] tuples for each measurement
//
// The function must:
// 1. Calculate apparent m/z values from flight times
// 2. Handle both 1+ and 2+ charged ions (2+ ions appear at half the expected m/z)
// 3. Group measurements by actual isotope
// 4. Return the weighted average atomic mass
//
// Returns: relative atomic mass rounded to ±0.01

function calculateRAM(E, d, data) {
  // Placeholder - solution below
}