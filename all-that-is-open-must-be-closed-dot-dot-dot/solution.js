function isBalanced(s, caps) {
  const openers = {};
  const closers = {};
  const selfPairs = new Set();

  for (let i = 0; i < caps.length; i += 2) {
    const open = caps[i];
    const close = caps[i + 1];
    if (open === close) {
      selfPairs.add(open);
    } else {
      openers[open] = close;
      closers[close] = open;
    }
  }

  const stack = [];
  for (const ch of s) {
    if (selfPairs.has(ch)) {
      if (stack.length && stack[stack.length - 1] === ch) {
        stack.pop();
      } else {
        stack.push(ch);
      }
    } else if (ch in openers) {
      stack.push(ch);
    } else if (ch in closers) {
      if (!stack.length || stack[stack.length - 1] !== closers[ch]) {
        return false;
      }
      stack.pop();
    }
  }

  return stack.length === 0;
}