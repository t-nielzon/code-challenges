function popBalloons(s, n, k) {
  // represent each balloon as an object with type and hit state
  let balloons = [];
  for (let i = 0; i < s.length; i++) {
    const ch = s[i];
    if (i + 1 < s.length && s[i + 1] === '*') {
      // regular balloon already hit once (starred) — shouldn't appear in initial input
      // but handle just in case
      balloons.push({ type: ch, starred: true });
      i++;
    } else {
      const isUpper = ch === ch.toUpperCase();
      balloons.push({ type: ch, starred: false, fortified: isUpper });
    }
  }

  for (let step = 0; step < k; step++) {
    if (balloons.length === 0) break;

    // target the n rightmost balloons
    const startIdx = Math.max(0, balloons.length - n);
    const hitIndices = new Set();
    for (let i = startIdx; i < balloons.length; i++) {
      hitIndices.add(i);
    }

    // apply hits and build new balloon array
    const newBalloons = [];
    for (let i = 0; i < balloons.length; i++) {
      if (!hitIndices.has(i)) {
        newBalloons.push(balloons[i]);
        continue;
      }

      const b = balloons[i];
      const t = b.type;

      if (b.starred) {
        // second hit on a fortified-turned-regular balloon
        const lower = t.toLowerCase();
        if (lower === 'w') {
          // pops
        } else if (lower === 'b') {
          newBalloons.push({ type: 'W', starred: false, fortified: true });
        } else if (lower === 'g') {
          newBalloons.push({ type: 'W', starred: false, fortified: true });
          newBalloons.push({ type: 'W', starred: false, fortified: true });
        } else if (lower === 'r') {
          newBalloons.push({ type: 'B', starred: false, fortified: true });
          newBalloons.push({ type: 'B', starred: false, fortified: true });
        }
      } else if (b.fortified) {
        // first hit on fortified → starred regular
        newBalloons.push({ type: t.toLowerCase(), starred: true, fortified: false });
      } else {
        // regular balloon hit
        const lower = t.toLowerCase();
        if (lower === 'w') {
          // pops
        } else if (lower === 'b') {
          newBalloons.push({ type: 'w', starred: false, fortified: false });
        } else if (lower === 'g') {
          newBalloons.push({ type: 'w', starred: false, fortified: false });
          newBalloons.push({ type: 'w', starred: false, fortified: false });
        } else if (lower === 'r') {
          newBalloons.push({ type: 'b', starred: false, fortified: false });
          newBalloons.push({ type: 'b', starred: false, fortified: false });
        }
      }
    }

    balloons = newBalloons;
  }

  // serialize back to string
  return balloons.map(b => b.type + (b.starred ? '*' : '')).join('');
}