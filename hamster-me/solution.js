function hamsterMe(code, message) {
  const codeLetters = [...new Set(code.split(''))].sort();
  const mapping = {};

  for (let i = 0; i < codeLetters.length; i++) {
    const start = codeLetters[i].charCodeAt(0) - 97;
    const next = codeLetters[(i + 1) % codeLetters.length].charCodeAt(0) - 97;
    let count = (next - start + 26) % 26;
    if (count === 0) count = 26;

    for (let j = 0; j < count; j++) {
      const letter = String.fromCharCode(((start + j) % 26) + 97);
      mapping[letter] = codeLetters[i] + (j + 1);
    }
  }

  return message.split('').map(ch => mapping[ch]).join('');
}