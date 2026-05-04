function scoreHand(cards) {
  let score = 0;
  let aces = 0;
  for (const card of cards) {
    if (card === 'A') {
      aces++;
      score += 11;
    } else if (['J', 'Q', 'K'].includes(card)) {
      score += 10;
    } else {
      score += parseInt(card, 10);
    }
  }
  while (score > 21 && aces > 0) {
    score -= 10;
    aces--;
  }
  return score;
}