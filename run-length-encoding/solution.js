function runLengthEncoding(str) {
  const result = [];
  for (let i = 0; i < str.length; i++) {
    if (result.length > 0 && result[result.length - 1][1] === str[i]) {
      result[result.length - 1][0]++;
    } else {
      result.push([1, str[i]]);
    }
  }
  return result;
}