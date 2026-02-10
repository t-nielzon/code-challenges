function snooker(scores) {
  const frames = scores.split(';').map(f => f.trim()).filter(f => f.length > 0);
  const result = [0, 0];

  for (const frame of frames) {
    // extract the two players' points by splitting on the dash that separates them
    // the dash separator is between two score segments: "score1-score2"
    // score can look like "24", "79(72)", or "105(53,52)"
    // we need to find the '-' that separates player1 from player2
    // it's the '-' that is NOT inside parentheses
    let depth = 0;
    let splitIdx = -1;
    for (let i = 0; i < frame.length; i++) {
      if (frame[i] === '(') depth++;
      else if (frame[i] === ')') depth--;
      else if (frame[i] === '-' && depth === 0) {
        splitIdx = i;
        break;
      }
    }

    const p1Str = frame.substring(0, splitIdx).trim();
    const p2Str = frame.substring(splitIdx + 1).trim();

    // extract the leading number (points) from a score string like "79(72)" or "105(53,52)"
    const getPoints = (s) => parseInt(s.match(/^\d+/)[0], 10);

    const p1Points = getPoints(p1Str);
    const p2Points = getPoints(p2Str);

    if (p1Points > p2Points) {
      result[0]++;
    } else if (p2Points > p1Points) {
      result[1]++;
    }
  }

  return result;
}