function translate(str) {
  // mapping: each consonant/semivowel gets the vowel that precedes it in the alphabet
  // a -> vowel (unchanged)
  // b,c,d -> last preceding vowel is 'a'
  // e -> vowel (unchanged)
  // f,g,h -> last preceding vowel is 'e'
  // i -> vowel (unchanged)
  // j,k,l,m,n -> last preceding vowel is 'i'
  // o -> vowel (unchanged)
  // p,q,r,s,t -> last preceding vowel is 'o'
  // u -> vowel (unchanged)
  // v,w,x,y,z -> last preceding vowel is 'u'

  const vowels = new Set(['a', 'e', 'i', 'o', 'u']);

  const precedingVowel = {};
  const alphabet = 'abcdefghijklmnopqrstuvwxyz';
  let lastVowel = '';
  for (const ch of alphabet) {
    if (vowels.has(ch)) {
      lastVowel = ch;
    } else {
      precedingVowel[ch] = lastVowel;
    }
  }

  let result = '';
  for (const ch of str) {
    const lower = ch.toLowerCase();
    if (precedingVowel[lower]) {
      result += ch + precedingVowel[lower];
    } else {
      result += ch;
    }
  }

  return result;
}

module.exports = { translate };