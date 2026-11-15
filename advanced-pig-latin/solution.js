function pigLatin(str) {
  const vowels = new Set('aeiouAEIOU');
  
  return str.split(' ').map(word => {
    // extract trailing punctuation
    let punctuation = '';
    let lastCharIdx = word.length - 1;
    while (lastCharIdx >= 0 && /[,.!?:;]/.test(word[lastCharIdx])) {
      punctuation = word[lastCharIdx] + punctuation;
      lastCharIdx--;
    }
    const actualWord = word.slice(0, lastCharIdx + 1);
    
    if (!actualWord) return punctuation;
    
    // numbers stay unchanged
    if (/^\d+$/.test(actualWord)) {
      return actualWord + punctuation;
    }
    
    const firstIsUpper = /[A-Z]/.test(actualWord[0]);
    const lowerWord = actualWord.toLowerCase();
    
    // find position of first vowel
    let firstVowelPos = -1;
    for (let i = 0; i < lowerWord.length; i++) {
      if (vowels.has(lowerWord[i])) {
        firstVowelPos = i;
        break;
      }
    }
    
    let transformed;
    if (firstVowelPos === -1) {
      // no vowels: append "ay"
      transformed = lowerWord + 'ay';
    } else if (firstVowelPos === 0) {
      // starts with vowel: append "way"
      transformed = lowerWord + 'way';
    } else {
      // move consonants before first vowel to end, append "ay"
      const consonants = lowerWord.slice(0, firstVowelPos);
      const rest = lowerWord.slice(firstVowelPos);
      transformed = rest + consonants + 'ay';
    }
    
    // restore original casing on first letter
    if (firstIsUpper) {
      transformed = transformed[0].toUpperCase() + transformed.slice(1);
    }
    
    return transformed + punctuation;
  }).join(' ');
}