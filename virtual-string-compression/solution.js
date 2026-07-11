function rle(length, getChar) {
  let result = '';
  let pos = 0;
  
  while (pos < length) {
    const char = getChar(pos);
    let count = 1;
    
    while (pos + count < length && getChar(pos + count) === char) {
      count++;
    }
    
    result += count + char;
    pos += count;
  }
  
  return result;
}