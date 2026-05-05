function stringExpansion(s) {
  let count = 1;
  let result = '';
  for (const ch of s) {
    if (/\d/.test(ch)) {
      count = Number(ch);
    } else {
      result += ch.repeat(count);
    }
  }
  return result;
}