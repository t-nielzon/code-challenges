function whoWillWin(arr) {
  if (!arr.length) return "No one raced!";

  const parseNum = str => parseFloat(str.match(/-?[\d.]+/)[0]);

  const valid = arr
    .map(r => {
      const speed = parseNum(r.speed);
      const distance = parseNum(r.distanceToRun);
      return { name: r.name, speed, distance };
    })
    .filter(r => r.speed > 0 && r.distance > 0);

  if (!valid.length) return "Everyone was disqualified!";

  const times = valid.map(r => ({
    name: r.name,
    time: Math.round(r.distance / r.speed)
  }));

  times.sort((a, b) => a.time - b.time);

  const bestTime = times[0].time;
  const winners = times.filter(t => t.time === bestTime);

  if (winners.length === 1) {
    return `${winners[0].name} won the race in ${bestTime} second(s)!`;
  }

  return `${winners.length} people tied in ${bestTime} second(s)!`;
}

module.exports = whoWillWin;