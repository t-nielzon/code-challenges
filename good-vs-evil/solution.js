function goodVsEvil(good, evil) {
  const goodWorth = [1, 2, 3, 3, 4, 10];
  const evilWorth = [1, 2, 2, 2, 3, 5, 10];

  const goodTotal = good.split(' ').reduce((sum, n, i) => sum + Number(n) * goodWorth[i], 0);
  const evilTotal = evil.split(' ').reduce((sum, n, i) => sum + Number(n) * evilWorth[i], 0);

  if (goodTotal > evilTotal) return 'Battle Result: Good triumphs over Evil';
  if (evilTotal > goodTotal) return 'Battle Result: Evil eradicates all trace of Good';
  return 'Battle Result: No victor on this battle field';
}

module.exports = { goodVsEvil };