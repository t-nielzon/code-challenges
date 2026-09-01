function berlinClock(timeStr) {
  const [hours, minutes, seconds] = timeStr.split(':').map(Number);
  
  // Top light - Yellow if even seconds, Off if odd
  const topLight = seconds % 2 === 0 ? 'Y' : 'O';
  
  // First row - Red fields for 5 hours each (0-4 lights)
  const fiveHours = Math.floor(hours / 5);
  const firstRow = 'R'.repeat(fiveHours) + 'O'.repeat(4 - fiveHours);
  
  // Second row - Red fields for 1 hour each (0-4 lights)
  const oneHour = hours % 5;
  const secondRow = 'R'.repeat(oneHour) + 'O'.repeat(4 - oneHour);
  
  // Third row - 11 lights for 5 minutes each
  // Red at positions 2, 5, 8 (representing 15, 30, 45 minutes)
  const fiveMinutes = Math.floor(minutes / 5);
  let thirdRow = '';
  for (let i = 0; i < 11; i++) {
    if (i < fiveMinutes) {
      thirdRow += (i === 2 || i === 5 || i === 8) ? 'R' : 'Y';
    } else {
      thirdRow += 'O';
    }
  }
  
  // Fourth row - 4 Yellow fields for 1 minute each
  const oneMinute = minutes % 5;
  const fourthRow = 'Y'.repeat(oneMinute) + 'O'.repeat(4 - oneMinute);
  
  return [topLight, firstRow, secondRow, thirdRow, fourthRow].join('\n');
}