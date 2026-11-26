function removeParentheses(str) {
  let depth = 0;
  let result = '';
  
  for (let char of str) {
    if (char === '(') {
      depth++;
    } else if (char === ')') {
      depth--;
    } else if (depth === 0) {
      result += char;
    }
  }
  
  return result;
}