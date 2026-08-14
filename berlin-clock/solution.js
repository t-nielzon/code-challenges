function berlinClock(time) {
  const [h, m, s] = time.split(':').map(Number);
  
  // Top row: seconds (Y if even, O if odd)
  const topRow = s % 2 === 0 ? 'Y' : 'O';
  
  // First row: 5-hour blocks (4 fields total)
  const firstRow = 'R'.repeat(Math.floor(h / 5)) + 'O'.repeat(4 - Math.floor(h / 5));
  
  // Second row: 1-hour blocks (4 fields total)
  const secondRow = 'R'.repeat(h % 5) + 'O'.repeat(4 - (h % 5));
  
  // Third row: 5-minute blocks (11 fields total, red at positions 3, 7, 11)
  const thirdRowCount = Math.floor(m / 5);
  let thirdRow = '';
  for (let i = 0; i < 11; i++) {
    if (i < thirdRowCount) {
      thirdRow += (i === 2 || i === 6 || i === 10) ? 'R' : 'Y';
    } else {
      thirdRow += 'O';
    }
  }
  
  // Fourth row: 1-minute blocks (4 fields total)
  const fourthRow = 'Y'.repeat(m % 5) + 'O'.repeat(4 - (m % 5));
  
  return `${topRow}\n${firstRow}\n${secondRow}\n${thirdRow}\n${fourthRow}`;
}