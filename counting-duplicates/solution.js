function duplicateCount(text) {
  const lower = text.toLowerCase();
  const charCount = {};
  
  for (let char of lower) {
    charCount[char] = (charCount[char] || 0) + 1;
  }
  
  return Object.values(charCount).filter(count => count > 1).length;
}