function verticalHistogram(s) {
  const counts = {};
  for (const char of s) {
    if (char >= 'A' && char <= 'Z') {
      counts[char] = (counts[char] || 0) + 1;
    }
  }
  
  const letters = Object.keys(counts).sort();
  
  if (letters.length === 0) return '';
  
  const maxCount = Math.max(...Object.values(counts));
  
  const rows = [];
  for (let row = maxCount; row >= 1; row--) {
    const line = letters
      .map(letter => counts[letter] >= row ? '*' : ' ')
      .join(' ')
      .trimEnd();
    rows.push(line);
  }
  
  rows.push(letters.join(' '));
  
  return rows.join('\n');
}