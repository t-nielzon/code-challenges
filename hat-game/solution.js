function guessColour(guesses, hats) {
  const countRed = arr => arr.reduce((a, c) => a + (c === "Red" ? 1 : 0), 0);

  if (guesses.length === 0) {
    return countRed(hats) % 2 === 0 ? "Red" : "Blue";
  }

  const parity = guesses[0] === "Red" ? 0 : 1;
  const known = countRed(guesses.slice(1)) + countRed(hats);
  return ((parity - known) % 2 + 2) % 2 === 1 ? "Red" : "Blue";
}