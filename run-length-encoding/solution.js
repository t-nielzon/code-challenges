function runLengthEncoding(str) {
  if (str.length === 0) return [];
  
  const result = [];
  let count = 1;
  let currentChar = str[0];
  
  for (let i = 1; i < str.length; i++) {
    if (str[i] === currentChar) {
      count++;
    } else {
      result.push([count, currentChar]);
      currentChar = str[i];
      count = 1;
    }
  }
  
  result.push([count, currentChar]);
  
  return result;
}