function rot135(str) {
  return str.replace(/./g, (char) => {
    const code = char.charCodeAt(0);
    
    // Handle lowercase letters
    if (code >= 97 && code <= 122) {
      return String.fromCharCode(((code - 97 + 13) % 26) + 97);
    }
    
    // Handle uppercase letters
    if (code >= 65 && code <= 90) {
      return String.fromCharCode(((code - 65 + 13) % 26) + 65);
    }
    
    // Handle digits
    if (code >= 48 && code <= 57) {
      return String.fromCharCode(((code - 48 + 5) % 10) + 48);
    }
    
    // Leave everything else unchanged
    return char;
  });
}