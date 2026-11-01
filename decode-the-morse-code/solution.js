function decodeMorse(morseCode) {
  const words = morseCode.trim().split('   ');
  
  return words.map(word => {
    const chars = word.split(' ');
    return chars.map(char => MORSE_CODE[char]).join('');
  }).join(' ');
}