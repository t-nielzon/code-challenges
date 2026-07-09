function simpleStringExpansion(str) {
  if (str === '') return '';
  
  let result = '';
  let multiplier = 1;
  
  for (let char of str) {
    if (/\d/.test(char)) {
      multiplier = parseInt(char);
    } else {
      result += char.repeat(multiplier);
    }
  }
  
  return result;
}