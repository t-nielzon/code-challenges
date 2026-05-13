function lookSay(num) {
  const s = String(num);
  let result = '';
  let i = 0;
  while (i < s.length) {
    let j = i;
    while (j < s.length && s[j] === s[i]) j++;
    result += (j - i) + s[i];
    i = j;
  }
  return Number(result);
}