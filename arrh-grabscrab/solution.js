function grabscrab(anagram, dictionary) {
  const sorted = anagram.split('').sort().join('');
  return dictionary.filter(word => word.split('').sort().join('') === sorted);
}