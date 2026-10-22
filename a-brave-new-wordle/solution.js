function wordleClues(guess, solution) {
  const result = ['b', 'b', 'b', 'b', 'b'];
  const solutionChars = solution.split('');
  
  // first pass: mark all correct positions (green)
  for (let i = 0; i < 5; i++) {
    if (guess[i] === solution[i]) {
      result[i] = 'g';
      solutionChars[i] = null; // mark character as used
    }
  }
  
  // second pass: mark wrong positions (yellow) and not found (black)
  for (let i = 0; i < 5; i++) {
    if (result[i] !== 'g') {
      const letterIndex = solutionChars.indexOf(guess[i]);
      if (letterIndex !== -1) {
        result[i] = 'y';
        solutionChars[letterIndex] = null; // mark character as used
      }
    }
  }
  
  return result.join('');
}