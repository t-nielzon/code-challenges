function grabscrab(jumble, dictionary) {
  const sortedJumble = jumble.split('').sort().join('');
  
  return dictionary.filter(word => {
    const sortedWord = word.split('').sort().join('');
    return sortedWord === sortedJumble;
  });
}