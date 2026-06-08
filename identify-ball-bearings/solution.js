/*
 * Strategy: take i+1 bearings from box i (i = 0..n-1) and weigh them all at
 * once. If every bearing were 'normal', the total would be the baseline
 * sum(1..n) * 10. The deluxe box contributes 1 extra gram per bearing taken
 * from it, so the surplus over the baseline equals (deluxeIndex + 1). This
 * pinpoints the deluxe box with a single weighing.
 */

function identifyBB(bearings, weigh) {
  const n = bearings.length;

  // build the sample: 1 bearing from box 0, 2 from box 1, ... (i+1 from box i)
  const sample = [];
  for (let i = 0; i < n; i++) {
    for (let k = 0; k <= i; k++) {
      sample.push(bearings[i]);
    }
  }

  const total = weigh(...sample);

  // baseline weight if all bearings were 'normal' (10g each)
  const baseline = (n * (n + 1) / 2) * 10;

  // surplus reveals (deluxeIndex + 1)
  const deluxeIndex = total - baseline - 1;

  return bearings[deluxeIndex];
}