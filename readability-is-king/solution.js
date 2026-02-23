function fleschKincaid(text) {
  const sentences = text.split(/[.!?:;]+/).filter(s => s.trim().length > 0);
  const totalSentences = sentences.length;

  // strip hyphens, dashes, apostrophes, parentheses, ellipses
  const cleaned = text.replace(/['\-\u2014\u2013()…]/g, '');

  // extract words (sequences of letters)
  const words = cleaned.match(/[a-zA-Z]+/g) || [];
  const totalWords = words.length;

  // count syllables per word: groups of consecutive vowels count as one
  const totalSyllables = words.reduce((sum, word) => {
    const vowelGroups = word.match(/[aeiou]+/gi) || [];
    return sum + Math.max(1, vowelGroups.length);
  }, 0);

  const avgWordsPerSentence = totalWords / totalSentences;
  const avgSyllablesPerWord = totalSyllables / totalWords;

  const grade = 0.39 * avgWordsPerSentence + 11.8 * avgSyllablesPerWord - 15.59;

  return Math.round(grade * 100) / 100;
}

module.exports = { fleschKincaid };