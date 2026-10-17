function reverseInParentheses(s) {
  function process(str, depth) {
    if (depth % 2 === 0) {
      // even depth: process nested parentheses without reversing this level
      let result = '';
      let i = 0;
      while (i < str.length) {
        if (str[i] === '(') {
          let count = 1;
          let j = i + 1;
          while (count > 0) {
            if (str[j] === '(') count++;
            else if (str[j] === ')') count--;
            j++;
          }
          let inner = str.substring(i + 1, j - 1);
          let processed = process(inner, depth + 1);
          result += '(' + processed + ')';
          i = j;
        } else {
          result += str[i];
          i++;
        }
      }
      return result;
    } else {
      // odd depth: reverse blocks (parentheses groups and text segments)
      let blocks = [];
      let i = 0;
      while (i < str.length) {
        if (str[i] === '(') {
          let count = 1;
          let j = i + 1;
          while (count > 0) {
            if (str[j] === '(') count++;
            else if (str[j] === ')') count--;
            j++;
          }
          blocks.push({
            type: 'paren',
            content: str.substring(i, j)
          });
          i = j;
        } else {
          let j = i;
          while (j < str.length && str[j] !== '(') j++;
          blocks.push({
            type: 'text',
            content: str.substring(i, j)
          });
          i = j;
        }
      }
      
      blocks.reverse();
      let result = '';
      for (let block of blocks) {
        if (block.type === 'paren') {
          let inner = block.content.substring(1, block.content.length - 1);
          let processed = process(inner, depth + 1);
          result += '(' + processed + ')';
        } else {
          let reversed = '';
          for (let k = block.content.length - 1; k >= 0; k--) {
            reversed += block.content[k];
          }
          result += reversed;
        }
      }
      return result;
    }
  }
  
  return process(s, 0);
}