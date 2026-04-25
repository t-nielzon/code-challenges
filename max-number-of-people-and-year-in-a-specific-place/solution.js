function yearMaxPeople(people) {
  const events = [];
  for (const [b, d] of people) {
    events.push([b, 1]);
    events.push([d + 1, -1]);
  }
  events.sort((a, b) => a[0] - b[0]);

  let count = 0, maxCount = 0, maxYear = events[0][0];
  let i = 0;
  while (i < events.length) {
    const year = events[i][0];
    while (i < events.length && events[i][0] === year) {
      count += events[i][1];
      i++;
    }
    if (count > maxCount) {
      maxCount = count;
      maxYear = year;
    }
  }
  return [maxCount, maxYear];
}