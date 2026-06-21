/*
 * Identify Ball Bearings — solution
 *
 * Classic counterfeit-coin / weighing puzzle.
 *
 * Take i+1 bearings of each box i (1 from box 0, 2 from box 1, ...) and weigh
 * them all in a single call. If every bearing were 'normal' the total would be
 * the baseline sum of counts * 10. The actual total exceeds the baseline by
 * exactly the count of deluxe bearings taken, which equals (deluxeIndex + 1).
 * That surplus pinpoints the deluxe box.
 */

function identifyBB(bearings, weigh) {
  // single box: it must be the deluxe one
  if (bearings.length === 1) return bearings[0];

  // build the sample: (i+1) copies of box i, all weighed at once
  const sample = [];
  let baselineCount = 0;
  for (let i = 0; i < bearings.length; i++) {
    for (let k = 0; k <= i; k++) sample.push(bearings[i]);
    baselineCount += i + 1;
  }

  const total = weigh(...sample);

  // surplus over an all-normal baseline equals (deluxeIndex + 1)
  const surplus = total - baselineCount * 10;
  return bearings[surplus - 1];
}