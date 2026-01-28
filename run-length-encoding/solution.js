function runLengthEncoding(str) {
  if (!str) return [];
  const result = [];
  let count = 1;
  for (let i = 1; i <= str.length; i++) {
    if (i < str.length && str[i] === str[i - 1]) {
      count++;
    } else {
      result.push([count, str[i - 1]]);
      count = 1;
    }
  }
  return result;
}