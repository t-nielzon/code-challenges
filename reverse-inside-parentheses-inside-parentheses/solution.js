function reverseInParentheses(str) {
  const chars = str.split('');
  const depths = new Array(chars.length).fill(0);
  const stack = [];

  for (let i = 0; i < chars.length; i++) {
    if (chars[i] === '(') {
      stack.push(i);
      depths[i] = stack.length;
    } else if (chars[i] === ')') {
      depths[i] = stack.length;
      stack.pop();
    }
  }

  function reverseRange(arr, l, r) {
    while (l < r) {
      [arr[l], arr[r]] = [arr[r], arr[l]];
      l++;
      r--;
    }
  }

  const openStack = [];
  for (let i = 0; i < chars.length; i++) {
    if (chars[i] === '(') {
      openStack.push(i);
    } else if (chars[i] === ')') {
      const start = openStack.pop();
      reverseRange(chars, start, i);
    }
  }

  for (let i = 0; i < chars.length; i++) {
    if (chars[i] === '(' || chars[i] === ')') {
      const depth = depths[i] !== 0 ? depths[i] : 0;
    }
  }

  for (let i = 0; i < chars.length; i++) {
    if (chars[i] === '(') chars[i] = '\x00';
    else if (chars[i] === ')') chars[i] = '\x01';
  }
  for (let i = 0; i < chars.length; i++) {
    if (chars[i] === '\x00') chars[i] = '(';
    else if (chars[i] === '\x01') chars[i] = ')';
  }

  return chars.join('');
}