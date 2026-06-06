const NIBBLES = (() => {
  const t = {};
  for (let i = 0; i < 16; i++) t[i.toString(2).padStart(4, '0')] = i.toString(16);
  return t;
})();

function binToHex(binary) {
  const groups = binary.match(/..../g);
  if (!groups) return '';
  let out = '';
  for (let i = 0; i < groups.length; i++) out += NIBBLES[groups[i]];
  return out;
}