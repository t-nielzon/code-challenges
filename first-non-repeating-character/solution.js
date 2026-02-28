function firstNonRepeatingLetter(s) {
  const chars = [...s];
  const counts = new Map();

  for (const ch of chars) {
    const lower = ch.toLowerCase();
    counts.set(lower, (counts.get(lower) || 0) + 1);
  }

  for (const ch of chars) {
    if (counts.get(ch.toLowerCase()) === 1) return ch;
  }

  return "";
}

module.exports = firstNonRepeatingLetter;