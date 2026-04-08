function pingPong(sounds) {
  const tokens = sounds.split('-');
  let pingScore = 0;
  let pongScore = 0;
  let i = 0;

  while (i < tokens.length) {
    const server = tokens[i];
    let lastHit = tokens[i];
    let j = i;

    while (j < tokens.length && (tokens[j] === 'ping' || tokens[j] === 'pong')) {
      lastHit = tokens[j];
      j++;
    }

    while (j < tokens.length && tokens[j] !== 'ping' && tokens[j] !== 'pong') {
      j++;
    }

    const winner = lastHit === 'ping' ? 'pong' : 'ping';
    if (server === winner) {
      if (winner === 'ping') pingScore++;
      else pongScore++;
    }

    i = j;
  }

  if (pingScore !== pongScore) return pingScore > pongScore ? 'ping' : 'pong';

  for (let k = tokens.length - 1; k >= 0; k--) {
    if (tokens[k] === 'ping' || tokens[k] === 'pong') {
      return tokens[k] === 'ping' ? 'pong' : 'ping';
    }
  }
}