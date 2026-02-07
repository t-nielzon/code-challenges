function whoWillWin(arr) {
  if (!arr.length) return "No one raced!";

  const qualified = arr
    .map(racer => {
      const speed = parseFloat(racer.speed);
      const distance = parseFloat(racer.distanceToRun);
      return { name: racer.name, speed, distance };
    })
    .filter(r => r.speed > 0 && r.distance > 0);

  if (!qualified.length) return "Everyone was disqualified!";

  const times = qualified.map(r => ({
    name: r.name,
    time: Math.round(r.distance / r.speed)
  }));

  const bestTime = Math.min(...times.map(t => t.time));
  const winners = times.filter(t => t.time === bestTime);

  if (winners.length === 1) {
    return `${winners[0].name} won the race in ${bestTime} second(s)!`;
  }

  return `${winners.length} people tied in ${bestTime} second(s)!`;
}