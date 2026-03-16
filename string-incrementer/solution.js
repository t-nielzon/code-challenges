function incrementString(strng) {
  const match = strng.match(/(.*?)(\d+)$/);
  if (!match) return strng + '1';
  const [, prefix, numStr] = match;
  const incremented = (BigInt(numStr) + 1n).toString();
  const padded = incremented.padStart(numStr.length, '0');
  return prefix + padded;
}

module.exports = { incrementString };