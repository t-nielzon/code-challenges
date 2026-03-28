function tongues(code) {
  const vowels = 'aiyeou';
  const consonants = 'bkxznhdcwgpvjqtsrlmf';

  function translate(ch) {
    const lower = ch.toLowerCase();
    let vi = vowels.indexOf(lower);
    if (vi !== -1) {
      const replaced = vowels[(vi + 3) % 6];
      return ch === lower ? replaced : replaced.toUpperCase();
    }
    let ci = consonants.indexOf(lower);
    if (ci !== -1) {
      const replaced = consonants[(ci + 10) % 20];
      return ch === lower ? replaced : replaced.toUpperCase();
    }
    return ch;
  }

  return [...code].map(translate).join('');
}

module.exports = { tongues };