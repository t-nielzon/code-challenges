function reverseOrRotate(str, sz) {
  if (sz <= 0 || str === "") {
    return "";
  }
  
  if (sz > str.length) {
    return "";
  }
  
  const chunks = [];
  for (let i = 0; i < str.length; i += sz) {
    const chunk = str.substring(i, i + sz);
    if (chunk.length < sz) {
      break;
    }
    chunks.push(chunk);
  }
  
  return chunks.map(chunk => {
    const sum = chunk.split('').reduce((acc, digit) => acc + parseInt(digit), 0);
    
    if (sum % 2 === 0) {
      return chunk.split('').reverse().join('');
    } else {
      return chunk.substring(1) + chunk[0];
    }
  }).join('');
}