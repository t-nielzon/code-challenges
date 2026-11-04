function findMissingLetter(array) {
  for (let i = 0; i < array.length - 1; i++) {
    const currentCode = array[i].charCodeAt(0);
    const nextCode = array[i + 1].charCodeAt(0);
    
    if (nextCode - currentCode > 1) {
      return String.fromCharCode(currentCode + 1);
    }
  }
}