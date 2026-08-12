function encode(n) {
  if (n === 0) return [0x00];
  
  const bytes = [];
  while (n > 0) {
    bytes.unshift(n & 0x7F);
    n >>= 7;
  }
  
  for (let i = 0; i < bytes.length - 1; i++) {
    bytes[i] |= 0x80;
  }
  
  return bytes;
}

function decode(data) {
  let value = 0;
  for (let i = 0; i < data.length; i++) {
    const byte = data[i];
    value = (value << 7) | (byte & 0x7F);
    if ((byte & 0x80) === 0) {
      break;
    }
  }
  return value;
}