function anagramsGame(subjects, memories, players) {
  const order = [
    [0, 2], [2, 1], [1, 3], [3, 0],
    [0, 3], [2, 0], [1, 2], [3, 1]
  ];
  const start = order.findIndex(p => p[0] === players[0] && p[1] === players[1]);
  const sets = memories.map(m => new Set(m));
  const sortKey = w => w.split('').sort().join('');
  const team = p => (p === 0 || p === 1) ? 0 : 1;

  const finals = [0, 0];

  for (let i = 0; i < subjects.length; i++) {
    const subject = subjects[i];
    const [p1, p2] = order[(start + i) % order.length];
    const key = sortKey(subject);

    let s1 = sets[p1].has(subject) ? 1 : 0;
    let s2 = sets[p2].has(subject) ? 1 : 0;

    const p1Anagrams = new Set();
    for (const w of sets[p1]) {
      if (w !== subject && sortKey(w) === key) {
        p1Anagrams.add(w);
        s1 += 2;
      }
    }
    for (const w of sets[p2]) {
      if (w !== subject && sortKey(w) === key && !p1Anagrams.has(w)) {
        s2 += 3;
      }
    }

    if (s1 > s2) finals[team(p1)]++;
    else if (s2 > s1) finals[team(p2)]++;
  }

  if (finals[0] > finals[1]) return 0;
  if (finals[1] > finals[0]) return 1;
  return -1;
}