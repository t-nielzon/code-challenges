function yearMaxPeople(people) {
  // a person present from birth..death counts through the death year, so their
  // leaving effect lands the year after; +1 at birth, -1 at death + 1
  const deltas = new Map();
  for (const [birth, death] of people) {
    deltas.set(birth, (deltas.get(birth) || 0) + 1);
    deltas.set(death + 1, (deltas.get(death + 1) || 0) - 1);
  }

  const years = [...deltas.keys()].sort((a, b) => a - b);

  let running = 0;
  let maxCount = -Infinity;
  let maxYear = years[0];

  for (const year of years) {
    running += deltas.get(year);
    if (running > maxCount) {
      maxCount = running;
      maxYear = year;
    }
  }

  return [maxCount, maxYear];
}