function fleschKincaid(text) {
  // split into sentences on . ! ?
  const sentences = text.split(/[.!?]+/).filter(s => s.trim().length > 0);
  const numSentences = sentences.length;

  // extract words: strip hyphens, dashes, apostrophes, parentheses, ellipses
  // then split on whitespace, keeping only words with letters
  const cleaned = text.replace(/['\-\(\)…]/g, '');
  const words = cleaned.split(/\s+/).filter(w => /[a-zA-Z]/.test(w));
  // remove trailing punctuation from each word
  const cleanWords = words.map(w => w.replace(/[^a-zA-Z]/g, ''));
  const numWords = cleanWords.length;

  // count syllables per word: groups of consecutive vowels (not y)
  function countSyllables(word) {
    const matches = word.toLowerCase().match(/[aeiou]+/g);
    return matches ? matches.length : 1;
  }

  const totalSyllables = cleanWords.reduce((sum, w) => sum + countSyllables(w), 0);

  const avgWordsPerSentence = numWords / numSentences;
  const avgSyllablesPerWord = totalSyllables / numWords;

  const grade = (0.39 * avgWordsPerSentence) + (11.8 * avgSyllablesPerWord) - 15.59;

  return Math.round(grade * 100) / 100;
}

module.exports = fleschKincaid;