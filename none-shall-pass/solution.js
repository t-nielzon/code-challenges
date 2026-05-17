function blackKnight(strikes) {
  const responses = [
    "'Tis but a scratch.",
    "Just a flesh wound!",
    "I'm invincible!",
    "All right, we'll call it a draw."
  ];
  const result = ["None shall pass."];
  let limbsLost = 0;
  for (const strike of strikes) {
    if (strike > 0 && limbsLost < 4) {
      result.push(responses[limbsLost]);
      limbsLost++;
    }
  }
  return result;
}