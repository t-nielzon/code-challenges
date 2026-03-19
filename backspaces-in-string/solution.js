function cleanString(s) {
  const result = [];
  for (const ch of s) {
    if (ch === '#') result.pop();
    else result.push(ch);
  }
  return result.join('');
}

module.exports = { cleanString };