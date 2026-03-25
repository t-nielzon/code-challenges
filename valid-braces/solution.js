function validBraces(braces) {
  const stack = [];
  const pairs = { '(': ')', '[': ']', '{': '}' };
  for (const ch of braces) {
    if (pairs[ch]) {
      stack.push(pairs[ch]);
    } else if (stack.pop() !== ch) {
      return false;
    }
  }
  return stack.length === 0;
}