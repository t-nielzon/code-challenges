function berlinClock(time) {
  const [hours, minutes, seconds] = time.split(':').map(Number);
  
  const topLight = seconds % 2 === 0 ? 'Y' : 'O';
  
  const fiveHourBlocks = Math.floor(hours / 5);
  const row1 = 'R'.repeat(fiveHourBlocks) + 'O'.repeat(4 - fiveHourBlocks);
  
  const oneHourBlocks = hours % 5;
  const row2 = 'R'.repeat(oneHourBlocks) + 'O'.repeat(4 - oneHourBlocks);
  
  const fiveMinuteBlocks = Math.floor(minutes / 5);
  let row3 = '';
  for (let i = 0; i < 11; i++) {
    if (i < fiveMinuteBlocks) {
      row3 += (i === 2 || i === 5 || i === 8) ? 'R' : 'Y';
    } else {
      row3 += 'O';
    }
  }
  
  const oneMinuteBlocks = minutes % 5;
  const row4 = 'Y'.repeat(oneMinuteBlocks) + 'O'.repeat(4 - oneMinuteBlocks);
  
  return `${topLight}\n${row1}\n${row2}\n${row3}\n${row4}`;
}