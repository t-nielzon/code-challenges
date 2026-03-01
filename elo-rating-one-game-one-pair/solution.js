function elo(experience, opponent, score, k) {
  const playerRating = experience.length === 0 ? 1000 : experience[experience.length - 1];

  const defaultK = (exp) => {
    if (exp.length < 30) return 25;
    if (exp.every(r => r < 2400)) return 15;
    return 10;
  };

  const kFn = k || defaultK;
  const expectation = 1 / (1 + Math.pow(10, (opponent - playerRating) / 400));
  const newRating = playerRating + kFn(experience) * (score - expectation);

  return Math.round(newRating);
}