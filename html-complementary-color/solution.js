function complementaryColor(hex) {
  if (typeof hex !== 'string') throw new Error('Invalid input');
  if (hex.length > 6) throw new Error('Invalid input');
  if (!/^[0-9a-fA-F]*$/.test(hex)) throw new Error('Invalid input');

  const padded = hex.padStart(6, '0');
  const num = parseInt(padded, 16);
  const comp = 0xFFFFFF - num;

  return '#' + comp.toString(16).toUpperCase().padStart(6, '0');
}

module.exports = { complementaryColor };