function incrementString(str) {
  const match = str.match(/^(.*?)(\d*)$/);
  const prefix = match[1];
  const numPart = match[2];

  if (numPart === '') {
    return prefix + '1';
  }

  const incremented = (BigInt(numPart) + 1n).toString();
  const padded = incremented.padStart(numPart.length, '0');
  return prefix + padded;
}