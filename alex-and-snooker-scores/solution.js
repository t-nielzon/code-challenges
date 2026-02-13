function snooker(scores) {
  const frames = scores.split(";").map(f => f.trim()).filter(f => f.length > 0);
  let p1Wins = 0;
  let p2Wins = 0;

  for (const frame of frames) {
    // extract the two players' points by splitting on "-"
    // each side may have parenthetical info we ignore for scoring
    const parts = frame.split("-");
    const p1Score = parseInt(parts[0], 10);
    const p2Score = parseInt(parts[1], 10);

    if (p1Score > p2Score) {
      p1Wins++;
    } else if (p2Score > p1Score) {
      p2Wins++;
    }
  }

  return [p1Wins, p2Wins];
}