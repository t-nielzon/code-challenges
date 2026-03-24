function pingPong(sounds) {
  const hits = sounds.split('-');
  let scores = { ping: 0, pong: 0 };
  let lastBadPlayer = null;

  let i = 0;
  while (i < hits.length) {
    // first hit of a rally is the serve
    const server = hits[i];
    let lastGoodHit = server;

    // advance through the rally
    let j = i + 1;
    while (j < hits.length && (hits[j] === 'ping' || hits[j] === 'pong')) {
      lastGoodHit = hits[j];
      j++;
    }

    // skip over bad noises until we reach the next serve or end
    while (j < hits.length && hits[j] !== 'ping' && hits[j] !== 'pong') {
      j++;
    }

    // lastGoodHit made the bad shot; the other player wins the rally
    // but only scores if that other player is also the server
    const badPlayer = lastGoodHit;
    lastBadPlayer = badPlayer;
    const winner = badPlayer === 'ping' ? 'pong' : 'ping';

    if (winner === server) {
      scores[server]++;
    }

    i = j;
  }

  if (scores.ping !== scores.pong) {
    return scores.ping > scores.pong ? 'ping' : 'pong';
  }

  // tied: winner is whoever did NOT hit the final bad shot
  return lastBadPlayer === 'ping' ? 'pong' : 'ping';
}

module.exports = { pingPong };