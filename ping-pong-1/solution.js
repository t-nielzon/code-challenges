function pingPong(sounds) {
  const hits = sounds.split('-');
  let server = null;
  let lastGoodHit = null;
  let inBadNoises = false;
  let finalBadShotBy = null;
  const scores = { ping: 0, pong: 0 };

  for (const hit of hits) {
    if (hit === 'ping' || hit === 'pong') {
      if (inBadNoises || server === null) {
        server = hit;
        inBadNoises = false;
      }
      lastGoodHit = hit;
    } else {
      if (!inBadNoises) {
        const badShotBy = lastGoodHit;
        const winner = badShotBy === 'ping' ? 'pong' : 'ping';
        if (winner === server) scores[winner]++;
        finalBadShotBy = badShotBy;
      }
      inBadNoises = true;
    }
  }

  if (scores.ping !== scores.pong) return scores.ping > scores.pong ? 'ping' : 'pong';
  return finalBadShotBy === 'ping' ? 'pong' : 'ping';
}