function stringExpansion(s) {
  const stack = [{ multiplier: 1, content: '' }];

  for (let i = 0; i < s.length; i++) {
    const ch = s[i];

    if (ch >= '1' && ch <= '9' && i + 1 < s.length && s[i + 1] === '(') {
      stack.push({ multiplier: Number(ch), content: '' });
      i++; // skip the '('
    } else if (ch === '(') {
      stack.push({ multiplier: 1, content: '' });
    } else if (ch === ')') {
      const { multiplier, content } = stack.pop();
      stack[stack.length - 1].content += content.repeat(multiplier);
    } else {
      stack[stack.length - 1].content += ch;
    }
  }

  return stack[0].content;
}

module.exports = { stringExpansion };