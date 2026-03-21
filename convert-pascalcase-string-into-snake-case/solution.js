function toUnderscore(string) {
  if (typeof string !== 'string') return String(string);
  return string.replace(/([A-Z])/g, (match, p1, offset) =>
    (offset ? '_' : '') + p1.toLowerCase()
  );
}

module.exports = toUnderscore;