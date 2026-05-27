function yearMaxPeople(people) {
  // a person present from arrival through death year inclusive: +1 at arrival,
  // -1 the year after death so the death year still counts them as present.
  const events = [];
  for (let i = 0; i < people.length; i++) {
    events.push([people[i][0], 1]);
    events.push([people[i][1] + 1, -1]);
  }

  // sort by year; within a year apply arrivals (+1) before departures (-1)
  // so the same-year peak is captured before anyone leaves.
  events.sort((a, b) => a[0] - b[0] || b[1] - a[1]);

  let count = 0;
  let maxCount = 0;
  let maxYear = events[0][0];

  for (let i = 0; i < events.length; i++) {
    count += events[i][1];
    if (count > maxCount) {
      maxCount = count;
      maxYear = events[i][0];
    }
  }

  return [maxCount, maxYear];
}