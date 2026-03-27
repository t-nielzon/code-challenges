function stringExpansion(s) {
  const stack = [{ mult: 1, str: '' }];

  for (let i = 0; i < s.length; i++) {
    const ch = s[i];
    if (ch >= '1' && ch <= '9' && s[i + 1] === '(') {
      stack.push({ mult: +ch, str: '' });
      i++; // skip '('
    } else if (ch === '(') {
      stack.push({ mult: 1, str: '' });
    } else if (ch === ')') {
      const { mult, str } = stack.pop();
      stack[stack.length - 1].str += str.repeat(mult);
    } else {
      stack[stack.length - 1].str += ch;
    }
  }

  return stack[0].str;
}

module.exports = { stringExpansion };