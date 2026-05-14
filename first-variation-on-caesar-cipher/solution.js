function shiftChar(ch, shift) {
  const code = ch.charCodeAt(0);
  if (code >= 65 && code <= 90) {
    return String.fromCharCode(((code - 65 + shift) % 26 + 26) % 26 + 65);
  }
  if (code >= 97 && code <= 122) {
    return String.fromCharCode(((code - 97 + shift) % 26 + 26) % 26 + 97);
  }
  return ch;
}

function movingShift(s, shift) {
  const encoded = s.split('').map((ch, i) => shiftChar(ch, shift + i)).join('');
  const len = encoded.length;
  const base = Math.floor(len / 5);
  const extra = len % 5;
  const parts = [];
  let idx = 0;
  for (let i = 0; i < 5; i++) {
    const size = base + (i < extra ? 1 : 0);
    parts.push(encoded.slice(idx, idx + size));
    idx += size;
  }
  return parts;
}

function demovingShift(s, shift) {
  const joined = s.join('');
  return joined.split('').map((ch, i) => shiftChar(ch, -(shift + i))).join('');
}