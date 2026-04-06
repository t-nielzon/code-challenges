function vowelRecognition(str) {
  const vowels = new Set('aeiouAEIOU');
  const n = str.length;
  let total = 0;
  for (let i = 0; i < n; i++) {
    if (vowels.has(str[i])) {
      total += (i + 1) * (n - i);
    }
  }
  return total;
}

module.exports = { vowelRecognition };