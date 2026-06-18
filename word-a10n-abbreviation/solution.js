function abbreviate(string) {
  return string.replace(/[a-zA-Z]+/g, (word) =>
    word.length >= 4
      ? `${word[0]}${word.length - 2}${word[word.length - 1]}`
      : word
  );
}