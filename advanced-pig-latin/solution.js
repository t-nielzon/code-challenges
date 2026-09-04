// solution.js
function pigLatin(str) {
  const words = str.split(' ');
  const vowels = new Set(['a', 'e', 'i', 'o', 'u']);
  
  const translated = words.map(word => {
    // Extract trailing punctuation
    let punctuation = '';
    let cleanWord = word;
    
    while (cleanWord.length > 0 && /[,.!?:;]/.test(cleanWord[cleanWord.length - 1])) {
      punctuation = cleanWord[cleanWord.length - 1] + punctuation;
      cleanWord = cleanWord.slice(0, -1);
    }
    
    // Numbers remain unchanged
    if (/^\d+$/.test(cleanWord)) {
      return word;
    }
    
    // Track if first letter is uppercase
    const isFirstCapital = cleanWord.length > 0 && /[A-Z]/.test(cleanWord[0]);
    
    // Work with lowercase version
    const lowerWord = cleanWord.toLowerCase();
    
    // Find first vowel
    let firstVowelIndex = -1;
    for (let i = 0; i < lowerWord.length; i++) {
      if (vowels.has(lowerWord[i])) {
        firstVowelIndex = i;
        break;
      }
    }
    
    let result;
    
    if (firstVowelIndex === -1) {
      // No vowels: append "ay"
      result = lowerWord + 'ay';
    } else if (firstVowelIndex === 0) {
      // First letter is a vowel: append "way"
      result = lowerWord + 'way';
    } else {
      // Move consonants before first vowel to end: add "ay"
      const consonants = lowerWord.slice(0, firstVowelIndex);
      const rest = lowerWord.slice(firstVowelIndex);
      result = rest + consonants + 'ay';
    }
    
    // Restore capitalization (only first letter)
    if (isFirstCapital) {
      result = result[0].toUpperCase() + result.slice(1);
    }
    
    return result + punctuation;
  });
  
  return translated.join(' ');
}