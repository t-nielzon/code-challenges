function maxMatch(sentence) {
  if (sentence.length === 0) return [];

  for (let end = sentence.length; end > 1; end--) {
    const word = sentence.substring(0, end);
    if (VALID_WORDS.has(word)) {
      return [word, ...maxMatch(sentence.substring(end))];
    }
  }

  return [sentence[0], ...maxMatch(sentence.substring(1))];
}