function solve(s) {
  let count = 0;
  for (let i = 0; i < s.length; i++) {
    if (parseInt(s[i]) % 2 !== 0) {
      count += i + 1;
    }
  }
  return count;
}