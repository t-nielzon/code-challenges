function reverseInParentheses(text) {
  let i = 0;

  function parse(depth) {
    const tokens = [];
    while (i < text.length && text[i] !== ')') {
      if (text[i] === '(') {
        i++; // consume '('
        const inner = parse(depth + 1);
        i++; // consume ')'
        tokens.push('(' + inner + ')');
      } else {
        tokens.push(text[i]);
        i++;
      }
    }
    // odd depths sit inside an odd number of parens, so they read reversed
    if (depth % 2 === 1) tokens.reverse();
    return tokens.join('');
  }

  return parse(0);
}