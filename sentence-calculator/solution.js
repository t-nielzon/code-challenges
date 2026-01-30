function calculateScore(sentence) {
  let score = 0;
  
  for (const char of sentence) {
    const code = char.charCodeAt(0);
    
    if (char >= 'a' && char <= 'z') {
      // Lower case: 'a'=1, 'b'=2, ..., 'z'=26
      score += code - 96;
    } else if (char >= 'A' && char <= 'Z') {
      // Upper case: 'A'=2, 'B'=4, ..., 'Z'=52
      score += (code - 64) * 2;
    } else if (char >= '0' && char <= '9') {
      // Digits: their numeric value
      score += code - 48;
    }
    // Other characters: 0 value (no addition needed)
  }
  
  return score;
}

module.exports = { calculateScore };