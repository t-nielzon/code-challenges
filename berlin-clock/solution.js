function berlinClock(time) {
  const [hours, minutes, seconds] = time.split(':').map(Number);
  
  const secondsLight = seconds % 2 === 0 ? 'Y' : 'O';
  
  const fiveHours = Math.floor(hours / 5);
  const firstRow = 'R'.repeat(fiveHours) + 'O'.repeat(4 - fiveHours);
  
  const oneHours = hours % 5;
  const secondRow = 'R'.repeat(oneHours) + 'O'.repeat(4 - oneHours);
  
  const fiveMinutes = Math.floor(minutes / 5);
  let thirdRow = '';
  for (let i = 0; i < 11; i++) {
    if (i < fiveMinutes) {
      thirdRow += (i === 2 || i === 5 || i === 8) ? 'R' : 'Y';
    } else {
      thirdRow += 'O';
    }
  }
  
  const oneMinutes = minutes % 5;
  const fourthRow = 'Y'.repeat(oneMinutes) + 'O'.repeat(4 - oneMinutes);
  
  return `${secondsLight}\n${firstRow}\n${secondRow}\n${thirdRow}\n${fourthRow}`;
}