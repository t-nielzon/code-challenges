function berlinClock(timeString) {
  const [hours, minutes, seconds] = timeString.split(':').map(Number);
  
  // Second indicator: Y for even, O for odd
  const secondRow = seconds % 2 === 0 ? 'Y' : 'O';
  
  // Five-hour indicator
  const fiveHours = Math.floor(hours / 5);
  const fiveHourRow = 'R'.repeat(fiveHours) + 'O'.repeat(4 - fiveHours);
  
  // One-hour indicator
  const oneHours = hours % 5;
  const oneHourRow = 'R'.repeat(oneHours) + 'O'.repeat(4 - oneHours);
  
  // Five-minute indicator
  const fiveMinutes = Math.floor(minutes / 5);
  let fiveMinuteRow = '';
  for (let i = 1; i <= 11; i++) {
    if (i <= fiveMinutes) {
      // Red at positions 3, 6, 9 (15, 30, 45 minutes)
      if (i === 3 || i === 6 || i === 9) {
        fiveMinuteRow += 'R';
      } else {
        fiveMinuteRow += 'Y';
      }
    } else {
      fiveMinuteRow += 'O';
    }
  }
  
  // One-minute indicator
  const oneMinutes = minutes % 5;
  const oneMinuteRow = 'Y'.repeat(oneMinutes) + 'O'.repeat(4 - oneMinutes);
  
  return [secondRow, fiveHourRow, oneHourRow, fiveMinuteRow, oneMinuteRow].join('\n');
}