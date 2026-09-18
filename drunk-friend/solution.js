function decode(str) {
  if (typeof str !== 'string') {
    return "Input is not a string";
  }
  
  return str.split('').map(char => {
    if (char >= 'a' && char <= 'z') {
      return String.fromCharCode('z'.charCodeAt(0) + 'a'.charCodeAt(0) - char.charCodeAt(0));
    } else if (char >= 'A' && char <= 'Z') {
      return String.fromCharCode('Z'.charCodeAt(0) + 'A'.charCodeAt(0) - char.charCodeAt(0));
    } else {
      return char;
    }
  }).join('');
}