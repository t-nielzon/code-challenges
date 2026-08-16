function guess_colour(guesses, hats) {
  if (guesses.length === 0) {
    // First player (at back): count red hats in front and signal parity
    // Odd number of reds → guess "Blue" (signal odd parity)
    // Even number of reds → guess "Red" (signal even parity)
    const redCount = hats.filter(h => h === "Red").length;
    return redCount % 2 === 1 ? "Blue" : "Red";
  } else {
    // Subsequent players: use parity signal to deduce own hat color
    const seenReds = hats.filter(h => h === "Red").length;
    const previousReds = guesses.filter(g => g === "Red").length;
    
    // If (visible reds + guessed reds so far) is odd → this player is blue
    // If even → this player is red
    return (seenReds + previousReds) % 2 === 1 ? "Blue" : "Red";
  }
}