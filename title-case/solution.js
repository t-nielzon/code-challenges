function titleCase(title, minorWords) {
  if (!title) return '';
  const minors = (minorWords || '').toLowerCase().split(' ');
  return title.toLowerCase().split(' ').map((word, i) => {
    if (i !== 0 && minors.includes(word)) return word;
    return word.charAt(0).toUpperCase() + word.slice(1);
  }).join(' ');
}