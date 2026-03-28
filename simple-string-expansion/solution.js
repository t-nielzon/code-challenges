function stringExpansion(s) {
  const stack = [{ str: '', multiplier: 1 }];

  let i = 0;
  while (i < s.length) {
    const ch = s[i];

    if (ch >= '1' && ch <= '9' && s[i + 1] === '(') {
      stack.push({ str: '', multiplier: Number(ch) });
      i += 2; // skip digit and '('
    } else if (ch === '(') {
      stack.push({ str: '', multiplier: 1 });
      i++;
    } else if (ch === ')') {
      const { str, multiplier } = stack.pop();
      stack[stack.length - 1].str += str.repeat(multiplier);
      i++;
    } else {
      stack[stack.length - 1].str += ch;
      i++;
    }
  }

  return stack[0].str;
}

module.exports = { stringExpansion };