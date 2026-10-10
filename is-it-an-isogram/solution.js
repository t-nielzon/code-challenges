function isIsogram(str) {
  if (str === '') return false;
  
  const letters = str.toLowerCase().replace(/[^a-z]/g, '');
  
  if (letters.length === 0) return false;
  
  const freq = {};
  for (const letter of letters) {
    freq[letter] = (freq[letter] || 0) + 1;
  }
  
  const frequencies = Object.values(freq);
  return frequencies.every(f => f === frequencies[0]);
}