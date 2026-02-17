function generateHashtag(str) {
  if (!str || !str.trim()) return false;

  const result = '#' + str
    .split(/\s+/)
    .filter(w => w.length > 0)
    .map(w => w[0].toUpperCase() + w.slice(1).toLowerCase())
    .join('');

  if (result.length > 140 || result === '#') return false;

  return result;
}

module.exports = { generateHashtag };